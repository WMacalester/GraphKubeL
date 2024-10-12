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
	InsertProductCategory(ctx context.Context, name string) (database.ProductCategory, error)
	GetProductCategories(ctx context.Context) ([]database.ProductCategory, error)
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

func (r *ProductRepository) InsertProductCategory(ctx context.Context, pc ProductCategory) (ProductCategory, error){
	productCategoryDao, err := r.Queries.InsertProductCategory(ctx, pc.Name)
	if err != nil {
		return ProductCategory{}, err
	}
	
	return ProductCategory{Id: int(productCategoryDao.ID), Name: productCategoryDao.Name}, nil
}

func (r *ProductRepository) GetProductCategories(ctx context.Context) ([]ProductCategory, error){
	daos, err := r.Queries.GetProductCategories(ctx);
	if err != nil {
		return nil, err
	}
	
	products := make([]ProductCategory, len(daos))
	for i, v := range daos {
		products[i] = mapProductCategoryDaoToProductCategory(v)
	}

	return products, nil
}

func mapProductDaoToProduct(dao database.Product) Product {
	return Product{Id: int(dao.ID), Name: dao.Name, Category: ProductCategory{} ,Description: dao.Description.String}
}

func mapProductCategoryDaoToProductCategory(dao database.ProductCategory) ProductCategory {
	return ProductCategory{Id: int(dao.ID), Name: dao.Name}
}
