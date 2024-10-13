package database

import (
	"context"
	"fmt"
	"os"

	"github.com/WMacalester/GraphKubeL/services/product/models"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductQueries interface {
	GetProducts(ctx context.Context) ([]GetProductsRow, error)
	GetProductCategoryById(ctx context.Context, id int32) (ProductCategory, error)
	GetProductCategories(ctx context.Context) ([]ProductCategory, error)
	InsertProduct(ctx context.Context, productCreateDao InsertProductParams) (Product, error)
	InsertProductCategory(ctx context.Context, name string) (ProductCategory, error)
}

type ProductRepository struct {
	Queries ProductQueries
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	queries := New(pool)
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

func (r *ProductRepository) GetProducts(ctx context.Context) ([]models.Product, error) {
	productDaos, err := r.Queries.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := make([]models.Product, len(productDaos))
	for i, v := range productDaos {
		products[i] = mapProductDaoToProduct(v)
	}

	return products, nil
}

func (r *ProductRepository) InsertProduct(ctx context.Context, product models.Product) (models.Product, error) {
	dao := InsertProductParams{
		Name: product.Name,
		CategoryID: pgtype.Int4{Int32: int32(product.Category.Id), Valid: true}, 
		Description: pgtype.Text{String: product.Description, Valid: true},
	}  

	val, err := r.Queries.InsertProduct(ctx, dao)
	if err != nil {
		return models.Product{}, err
	}

	return models.Product{Id: int(val.ID), Name: val.Name, Description: val.Description.String}, nil
}

func (r *ProductRepository) InsertProductCategory(ctx context.Context, pc models.ProductCategory) (models.ProductCategory, error){
	productCategoryDao, err := r.Queries.InsertProductCategory(ctx, pc.Name)
	if err != nil {
		return models.ProductCategory{}, err
	}
	
	return models.ProductCategory{Id: int(productCategoryDao.ID), Name: productCategoryDao.Name}, nil
}

func (r *ProductRepository) GetProductCategoryById(ctx context.Context, id int32) (models.ProductCategory, error) {
	dao, err := r.Queries.GetProductCategoryById(ctx, id)
	if err != nil {
		return models.ProductCategory{}, err
	}
	return mapProductCategoryDaoToProductCategory(dao), nil
}

func (r *ProductRepository) GetProductCategories(ctx context.Context) ([]models.ProductCategory, error){
	daos, err := r.Queries.GetProductCategories(ctx);
	if err != nil {
		return nil, err
	}
	
	products := make([]models.ProductCategory, len(daos))
	for i, v := range daos {
		products[i] = mapProductCategoryDaoToProductCategory(v)
	}

	return products, nil
}

func mapProductDaoToProduct(dao GetProductsRow) models.Product {
	return models.Product{Id: int(dao.ID), Name: dao.Name, Category: models.ProductCategory{Id: int(dao.CategoryID), Name: dao.CategoryName} ,Description: dao.Description.String}
}

func mapProductCategoryDaoToProductCategory(dao ProductCategory) models.ProductCategory {
	return models.ProductCategory{Id: int(dao.ID), Name: dao.Name}
}
