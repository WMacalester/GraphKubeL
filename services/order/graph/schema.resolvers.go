package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
)

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*Order, error) {
	orders, err := r.OrderRepository.GetOrders(ctx)
	if err != nil {
		return nil, err
	}


	orderDtos := make([]*Order, len(orders))

	for _, order := range orders {
		orderDto := mapOrderToDto(order)
		orderDtos = append(orderDtos, &orderDto)
	}

	return orderDtos, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }