package model

type Order struct {
	Products []Product
	Total float64
}

func NewOrder(products []Product, total float64) *Order {
	o := Order{Products: products, Total: total}
	return &o
}
