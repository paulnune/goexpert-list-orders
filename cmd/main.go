package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"goexpert-list-orders/internal/db"
	grpcService "goexpert-list-orders/internal/delivery/grpc"
	pb "goexpert-list-orders/internal/delivery/grpc/pb"
	"goexpert-list-orders/internal/delivery/rest"
	"goexpert-list-orders/internal/usecase"

	_ "github.com/lib/pq" // Driver PostgreSQL
	"google.golang.org/grpc"
)

func main() {
	// Configuração da string de conexão
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Conexão com o banco de dados
	dbConn, err := connectToDB(connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer dbConn.Close()

	// Inicialização do repositório, caso de uso e handlers
	repo := db.NewOrderRepository(dbConn)
	listOrdersUC := usecase.NewListOrdersUseCase(repo)
	restHandler := rest.NewHandler(listOrdersUC)
	grpcHandler := grpcService.NewServer(listOrdersUC)

	// Inicializar servidores
	go startRESTServer(restHandler)
	startGRPCServer(grpcHandler)
}

// Conecta ao banco de dados com tentativas de retry
func connectToDB(connStr string) (*sql.DB, error) {
	var dbConn *sql.DB
	var err error
	for retries := 0; retries < 10; retries++ {
		dbConn, err = sql.Open("postgres", connStr)
		if err == nil {
			err = dbConn.Ping()
			if err == nil {
				fmt.Println("Conexão com o banco de dados bem-sucedida!")
				return dbConn, nil
			}
		}
		fmt.Printf("Erro ao conectar: %v. Tentando novamente...\n", err)
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("não foi possível conectar ao banco de dados após várias tentativas: %v", err)
}

// Inicializa o servidor REST
func startRESTServer(handler *rest.Handler) {
	http.HandleFunc("/orders", handler.ListOrders)
	http.HandleFunc("/order", handler.CreateOrder)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	fmt.Println("Servidor REST rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Inicializa o servidor gRPC
func startGRPCServer(grpcHandler pb.OrderServiceServer) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao iniciar o servidor gRPC: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, grpcHandler)

	fmt.Println("Servidor gRPC rodando na porta 50051")
	log.Fatal(grpcServer.Serve(lis))
}
