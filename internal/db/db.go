package db

import (
	"context"
	"database/sql"
	"goexpert-list-orders/internal/domain"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll(ctx context.Context) ([]domain.Order, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, customer, total FROM orders")
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
