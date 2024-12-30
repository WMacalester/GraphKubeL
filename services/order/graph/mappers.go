package graph

import (
	"github.com/WMacalester/GraphKubeL/services/order/models"
)

func mapOrderToDto(order models.Order) Order {
	return Order{ID: order.Id, TransactionID: &order.TransactionID, ProductID: &order.ProductId, NumberOfItems: &order.NumberOfItems}
}