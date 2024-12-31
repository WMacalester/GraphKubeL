package database

import (
	"context"
	"os"

	"github.com/WMacalester/GraphKubeL/internal/common"
	"github.com/WMacalester/GraphKubeL/services/order/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderQueries interface {
	GetOrders(ctx context.Context) ([]Order, error)
	InsertOrder(ctx context.Context, orderDao InsertOrderParams) (Order, error)
}

type OrderRepository struct {
	Queries OrderQueries
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	queries := New(pool)
	return &OrderRepository{Queries: queries}
}

func CreateConnString() (string, error){
	user := os.Getenv("PGUSER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("ORDER_PG_HOST")
	port := os.Getenv("ORDER_DB_PORT")
	dbname := os.Getenv("POSTGRES_DB")
	sslmode := os.Getenv("ORDER_PG_SSLMODE")

	return common.FormatPostgresConnString(user, password, host, port, dbname, sslmode)
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

func (r *OrderRepository) InsertOrder(ctx context.Context, newOrder models.Order)(models.Order, error){
	insertOrderParamsDao := mapOrderToInsertOrderParamsDao(newOrder) 
	savedOrder, err := r.Queries.InsertOrder(ctx, insertOrderParamsDao)
	if err != nil {
		return models.Order{}, nil
	}
	return mapOrderDaoToOrder(savedOrder), nil
}

func mapOrderDaoToOrder(dao Order) models.Order {
	return models.Order{TransactionID: int(dao.TransactionID), ProductId: int(dao.ProductID), NumberOfItems: int(dao.NumberOfItems)}
}

func mapOrderToInsertOrderParamsDao(model models.Order) InsertOrderParams {
	return InsertOrderParams{TransactionID: int32(model.TransactionID), ProductID: int32(model.ProductId), NumberOfItems: int32(model.NumberOfItems)}
}
