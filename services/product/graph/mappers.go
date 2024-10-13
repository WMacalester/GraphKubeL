package graph

import (
	"github.com/WMacalester/GraphKubeL/services/product/graph/model"
	"github.com/WMacalester/GraphKubeL/services/product/models"
)

func mapProductCategoryToProductCategoryDto(pc models.ProductCategory) *model.ProductCategory { 
	return &model.ProductCategory{ID: pc.Id, Name: pc.Name}
}