// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

type Order struct {
	TransactionID int `json:"transactionId"`
	ProductID     int `json:"productId"`
	NumberOfItems int `json:"numberOfItems"`
}

func (Order) IsEntity() {}

type Query struct {
}
