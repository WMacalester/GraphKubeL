package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"

	"github.com/WMacalester/GraphKubeL/services/product/graph/model"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.ProductCreateDto) (*model.Product, error) {
	product := &model.Product{Name: &input.Name, Category: &input.Category, Description: &input.Description}
	r.products = append(r.products, product)
	return product, nil
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context) ([]*model.Product, error) {
	return r.products, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
