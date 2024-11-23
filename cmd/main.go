package main

import (
	"log"
	"net"
	"net/http"

	"goexpert-list-orders/configs"
	"goexpert-list-orders/internal/db"
	grpcDelivery "goexpert-list-orders/internal/delivery/grpc"
	"goexpert-list-orders/internal/delivery/rest"
	"goexpert-list-orders/internal/usecase"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func main() {
	// Configurações
	configs.LoadConfig()

	// Banco de dados
	conn, err := db.NewDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	// Repositório e caso de uso
	repo := db.NewOrderRepository(conn)
	listOrdersUC := &usecase.ListOrdersUseCase{Repo: repo}

	// Servidor REST
	go func() {
		router := mux.NewRouter()
		handler := &rest.Handler{ListOrdersUC: listOrdersUC}
		router.HandleFunc("/order", handler.ListOrders).Methods("GET")

		log.Println("Servidor REST iniciado na porta 8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	}()

	// Servidor gRPC
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Erro ao iniciar listener gRPC: %v", err)
		}

		grpcServer := grpc.NewServer()
		grpcService := &grpcDelivery.Server{ListOrdersUC: listOrdersUC}
		grpcDelivery.RegisterOrderServiceServer(grpcServer, grpcService)

		log.Println("Servidor gRPC iniciado na porta 50051")
		log.Fatal(grpcServer.Serve(lis))
	}()

	// Mantém o programa ativo
	select {}
}
