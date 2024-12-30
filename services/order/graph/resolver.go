package graph

import "github.com/WMacalester/GraphKubeL/services/order/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	OrderRepository *database.OrderRepository
}
