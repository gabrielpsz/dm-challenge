package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Quantity int `json:"quantity"`
	oldProductQuantity int
}

func (product Product) GetOldProductQuantity() int {
	return product.oldProductQuantity;
}

func (product *Product) SetOldProductQuantity(oldQuantity int) {
	product.oldProductQuantity = oldQuantity
}

func NewProduct(name string, price float64, quantity int) *Product {
	p := Product{Name: name, Price: price, Quantity: quantity}
	return &p
}
