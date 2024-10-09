package main

import (
	"context"
	"fmt"
	"os"

	"github.com/WMacalester/GraphKubeL/services/product/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductQueries interface {
	GetProducts(ctx context.Context) ([]database.Product, error)
}

type ProductRepository struct {
	Queries ProductQueries
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	queries := database.New(pool)
	return &ProductRepository{Queries: queries}
}

func CreateConnString() (string, error){
	user := os.Getenv("PRODUCT_PG_USER")
	password := os.Getenv("PRODUCT_PG_PASSWORD")
	host := os.Getenv("PRODUCT_PG_HOST")
	port := os.Getenv("PRODUCT_PG_PORT")
	dbname := os.Getenv("PRODUCT_PG_DATABASE")
	sslmode := os.Getenv("PRODUCT_PG_SSLMODE")

	if user == "" || password == "" || host == "" || port == "" || dbname == "" {
		return "", fmt.Errorf("one or more required environment variables are missing")
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode)

	return connString, nil
}

func (r *ProductRepository) GetProducts(ctx context.Context) ([]Product, error) {
	productDaos, err := r.Queries.GetProducts(ctx)

	if err != nil {
		return nil, err
	}

	products := make([]Product, len(productDaos))
	for i, v := range productDaos {
		products[i] = mapProductDaoToProduct(v)
	}

	return products, nil
}

func mapProductDaoToProduct(dao database.Product) Product {
	return Product{Name: dao.Name, Category: ProductCategory(dao.Category.Int32) ,Description: dao.Description.String}
}