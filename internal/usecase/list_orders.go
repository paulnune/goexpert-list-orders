package usecase

import (
	"context"
	"goexpert-list-orders/internal/db"
	"goexpert-list-orders/internal/domain"
)

type ListOrdersUseCase struct {
	repository *db.OrderRepository
}

func NewListOrdersUseCase(repository *db.OrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{repository: repository}
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context) ([]domain.Order, error) {
	return uc.repository.ListOrders(ctx)
}

func (uc *ListOrdersUseCase) CreateOrder(ctx context.Context, order domain.Order) (int, error) {
	return uc.repository.CreateOrder(ctx, order)
}
