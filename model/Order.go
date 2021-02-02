package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Products []Product `json:"products"`
	Total float64 `json:"total"`
}

func NewOrder(products []Product, total float64) *Order {
	o := Order{ID: primitive.NewObjectID(), Products: products, Total: total}
	return &o
}
