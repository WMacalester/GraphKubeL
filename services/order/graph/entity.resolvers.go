package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"
)

// FindOrderByTransactionIDAndProductID is the resolver for the findOrderByTransactionIDAndProductID field.
func (r *entityResolver) FindOrderByTransactionIDAndProductID(ctx context.Context, transactionID int, productID int) (*Order, error) {
	panic(fmt.Errorf("not implemented: FindOrderByTransactionIDAndProductID - findOrderByTransactionIDAndProductID"))
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
