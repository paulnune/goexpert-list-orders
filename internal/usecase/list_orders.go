package usecase

import (
	"context"
	"goexpert-list-orders/internal/domain"
)

// Define a interface para o repositório de ordens
type OrderRepository interface {
	ListOrders() ([]domain.Order, error)
}

// Caso de uso para listar ordens
type ListOrdersUseCase struct {
	Repo OrderRepository
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context) ([]domain.Order, error) {
	return uc.Repo.ListOrders(ctx)
}
