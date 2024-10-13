package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"

	"github.com/WMacalester/GraphKubeL/services/product/graph/model"
	"github.com/WMacalester/GraphKubeL/services/product/models"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.ProductCreateDto) (*model.Product, error) {
	product := models.Product{Name: input.Name, Category: models.ProductCategory{Id: input.CategoryID}, Description: input.Description}

	inserted, err := r.Repository.InsertProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return mapProductToProductDto(inserted), nil
}

// CreateProductCategory is the resolver for the createProductCategory field.
func (r *mutationResolver) CreateProductCategory(ctx context.Context, input *model.ProductCategoryCreateDto) (*model.ProductCategory, error) {
	pc := models.ProductCategory{Name: input.Name}
	saved, err := r.Repository.InsertProductCategory(ctx, pc)
	if err != nil {
		return nil, err
	}

	return mapProductCategoryToProductCategoryDto(saved), nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	products, err := r.Repository.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	dtos := make([]*model.Product, len(products))
	for i, val := range products {
		dto := mapProductToProductDto(val)
		dtos[i] = dto
	}

	return dtos, nil
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id int) (*model.Product, error) {
	product, err := r.Repository.GetProductById(ctx, int32(id))
	if err != nil {
		return &model.Product{}, err
	}

	return mapProductToProductDto(product), nil
}

// ProductCategories is the resolver for the ProductCategories field.
func (r *queryResolver) ProductCategories(ctx context.Context) ([]*model.ProductCategory, error) {
	pcs, err := r.Repository.GetProductCategories(ctx)
	if err != nil {
		return nil, err
	}

	pcsPtr := make([]*model.ProductCategory, len(pcs))
	for i, val := range pcs {
		dto := mapProductCategoryToProductCategoryDto(val)
		pcsPtr[i] = dto
	}

	return pcsPtr, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
