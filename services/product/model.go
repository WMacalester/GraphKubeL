package main

type Product struct {
	Name string
	Category ProductCategory 
	Description string
}

type ProductCategory int

const (
	Clothing ProductCategory = iota
	Food
)
