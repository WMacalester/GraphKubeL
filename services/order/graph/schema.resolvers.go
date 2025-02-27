package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"

	"github.com/WMacalester/GraphKubeL/services/order/models"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *OrderCreateDto) (*Order, error) {
	order := models.Order{TransactionID: input.TransactionID, ProductId: input.ProductID, NumberOfItems: input.NumberOfItems}

	insertedOrder, err := r.OrderRepository.InsertOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	dto := mapOrderToDto(insertedOrder)
	return &dto, nil
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*Order, error) {
	orders, err := r.OrderRepository.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	orderDtos := make([]*Order, 0, len(orders))

	for _, order := range orders {
		orderDto := mapOrderToDto(order)
		orderDtos = append(orderDtos, &orderDto)
	}

	return orderDtos, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
