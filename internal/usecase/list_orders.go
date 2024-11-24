package usecase

import (
	"context"
	"goexpert-list-orders/internal/domain"
)

type OrderRepository interface {
	ListOrders(ctx context.Context) ([]domain.Order, error)
}

type ListOrdersUseCase struct {
	Repo OrderRepository
}

func NewListOrdersUseCase(repo OrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{Repo: repo}
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context) ([]domain.Order, error) {
	orders, err := uc.Repo.ListOrders(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
