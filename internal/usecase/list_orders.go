package usecase

import (
	"context"
	"goexpert-list-orders/internal/domain"
)

type OrderRepository interface {
	FindAll(ctx context.Context) ([]domain.Order, error)
}

type ListOrdersUseCase struct {
	Repo OrderRepository
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context) ([]domain.Order, error) {
	orders, err := uc.Repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
