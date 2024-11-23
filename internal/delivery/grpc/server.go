package grpc

import (
	"context"
	"goexpert-list-orders/internal/delivery/grpc/pb"
	"goexpert-list-orders/internal/usecase"
)

type Server struct {
	pb.UnimplementedOrderServiceServer
	ListOrdersUC *usecase.ListOrdersUseCase
}

func (s *Server) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.ListOrdersUC.Execute(ctx)
	if err != nil {
		return nil, err
	}

	var grpcOrders []*pb.Order
	for _, order := range orders {
		grpcOrders = append(grpcOrders, &pb.Order{
			Id:       order.ID,
			Customer: order.Customer,
			Total:    order.Total,
		})
	}

	return &pb.ListOrdersResponse{Orders: grpcOrders}, nil
}
