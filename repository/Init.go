package repository

import (
	"context"
	"github.com/gabrielpsz/dm-challenge/model"
	"github.com/gabrielpsz/dm-challenge/tools"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

var ctx = context.TODO()
var productsCollection *mongo.Collection
var ordersCollection *mongo.Collection

func StartDatabase() {
	fmt.Println("Starting database.")
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database("dm-challenge")
	productsCollection = database.Collection("products")
	ordersCollection = database.Collection("orders")
	collections, err := database.ListCollectionNames(ctx, bson.D{{}})
	if (!contains(collections, "products")) {
		InsertProductData("files/products.csv")
	}
}

func InsertProductData(productFilePath string) {
	records, err := tools.ReadData(productFilePath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	if err != nil {
		log.Fatal(err)
	}
	records = records[1:]
	for _, record := range records {
		if price, err := strconv.ParseFloat(record[1], 64); err == nil {
			product := model.NewProduct(record[0], price, 0)
			InsertProduct(product)
		}
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}