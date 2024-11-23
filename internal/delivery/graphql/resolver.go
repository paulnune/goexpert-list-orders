package graphql

import (
	"context"
	"order-service/internal/usecase"
)

type Resolver struct {
	ListOrdersUC *usecase.ListOrdersUseCase
}

func (r *Resolver) ListOrders(ctx context.Context) ([]*Order, error) {
	orders, err := r.ListOrdersUC.Execute(ctx)
	if err != nil {
		return nil, err
	}

	var gqlOrders []*Order
	for _, o := range orders {
		gqlOrders = append(gqlOrders, &Order{
			ID:           o.ID,
			CustomerName: o.CustomerName,
			Amount:       o.Amount,
			CreatedAt:    o.CreatedAt.String(),
		})
	}

	return gqlOrders, nil
}
