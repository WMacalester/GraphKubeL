package graph

import (
	"github.com/WMacalester/GraphKubeL/services/product/database"
	"github.com/WMacalester/GraphKubeL/services/product/graph/model"
)

type Resolver struct{
	products []*model.Product
	Repository *database.ProductRepository
}
