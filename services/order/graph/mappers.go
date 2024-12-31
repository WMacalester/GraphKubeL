package graph

import (
	"github.com/WMacalester/GraphKubeL/services/order/models"
)

func mapOrderToDto(order models.Order) Order {
	return Order{TransactionID: order.TransactionID, ProductID: order.ProductId, NumberOfItems: order.NumberOfItems}
}