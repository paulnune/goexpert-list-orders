package main

import (
	"log"
	"net"
	"net/http"

	"goexpert-list-orders/configs"
	"goexpert-list-orders/internal/db"
	"goexpert-list-orders/internal/delivery/grpc"
	"goexpert-list-orders/internal/delivery/grpc/pb"
	"goexpert-list-orders/internal/delivery/rest"
	"goexpert-list-orders/internal/usecase"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	// Configurações iniciais
	configs.LoadConfig()

	// Inicializa banco de dados
	conn, err := db.NewDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	// Inicializa o caso de uso
	repo := db.NewRepository(conn)
	listOrdersUC := &usecase.ListOrdersUseCase{Repo: repo}

	// Inicia servidor REST
	go func() {
		router := mux.NewRouter()
		handler := &rest.Handler{ListOrdersUC: listOrdersUC}

		router.HandleFunc("/order", handler.ListOrders).Methods("GET")

		log.Println("Servidor REST iniciado na porta 8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	}()

	// Inicia servidor gRPC
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Erro ao iniciar listener gRPC: %v", err)
		}

		grpcServer := grpc.NewServer()
		grpcService := &grpc.Server{ListOrdersUC: listOrdersUC}

		pb.RegisterOrderServiceServer(grpcServer, grpcService)

		log.Println("Servidor gRPC iniciado na porta 50051")
		log.Fatal(grpcServer.Serve(lis))
	}()

	select {}
}
