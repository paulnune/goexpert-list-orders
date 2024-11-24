package db

import (
	"context"
	"database/sql"
	"goexpert-list-orders/internal/domain"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) ListOrders(ctx context.Context) ([]domain.Order, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, customer, total FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.ID, &order.Customer, &order.Total); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
