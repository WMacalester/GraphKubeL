package models

type Product struct {
	Id          int
	Name        string
	Category    ProductCategory
	Description string
}

type ProductCategory struct {
	Id   int
	Name string
}
