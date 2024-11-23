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

func (s *Server) ListOrders(ctx context.Context, req *pb.Empty) (*pb.OrderListResponse, error) {
	orders, err := s.ListOrdersUC.Execute()
	if err != nil {
		return nil, err
	}

	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			Id:       order.ID,
			Customer: order.Customer,
			Total:    float32(order.Total),
		})
	}

	return &pb.OrderListResponse{Orders: pbOrders}, nil
}
