package database

import (
	"context"

	"github.com/WMacalester/GraphKubeL/services/order/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderQueries interface {
	GetOrders(ctx context.Context) ([]Order, error)
}

type OrderRepository struct {
	Queries OrderQueries
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	queries := New(pool)
	return &OrderRepository{Queries: queries}
}

func (r *OrderRepository) GetOrders(ctx context.Context) ([]models.Order, error){
	daos, err := r.Queries.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	var orders []models.Order
	for _, dao := range daos {
		orders = append(orders, mapOrderDaoToOrder(dao))
	}

	return orders, nil
}

func mapOrderDaoToOrder(dao Order) models.Order {
	return models.Order{Id: int(dao.ID), TransactionID: int(dao.TransactionID), ProductId: int(dao.ProductID), NumberOfItems: int(dao.NumberOfItems)}
}