package repository

import (
	"context"
	"errors"
	"github.com/gabrielpsz/dm-challenge/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	ORDERS = "orders"
)

func InsertOrder(order *model.Order) error {

	insertResult, err := ordersCollection.InsertOne(ctx, order)

	var newProduct model.Product
	for _, product := range order.Products {
		newProduct.ID = product.ID
		newProduct.Name = product.Name
		newProduct.Price = product.Price
		newProduct.Quantity = product.GetOldProductQuantity() - product.Quantity
		if (newProduct.Quantity < 0) {
			message := fmt.Sprintf("A quantidade requisitada do produto %s está fora de estoque. Quantidade total disponível: %v", product.Name, product.GetOldProductQuantity())
			return errors.New(message)
		} else {
			UpdateProduct(product.ID.Hex(), &newProduct)
		}
	}

	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("Inserted order with ID:", insertResult.InsertedID)
	return err
}

func GetOrderById(id string) (model.Order) {
	var order model.Order
	var orderModel bson.M
	modelId, _ := primitive.ObjectIDFromHex(id)
	err := ordersCollection.FindOne(ctx, bson.M{"_id": modelId}).Decode(&orderModel)
	if err != nil { return order }
	bsonBytes, _ := bson.Marshal(orderModel)
	bson.Unmarshal(bsonBytes, &order)
	return order
}

func GetOrders() ([]model.Order) {
	var orders []model.Order
	cursor, err := ordersCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var order model.Order
		err := cursor.Decode(&order)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)

	}
	return orders
}
