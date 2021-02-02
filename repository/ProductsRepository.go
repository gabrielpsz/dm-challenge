package repository

import (
	"github.com/gabrielpsz/dm-challenge/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//mgobson "gopkg.in/mgo.v2/bson"
	"log"
)

const (
	PRODUCTS = "products"
)


func GetProductByName(name string) ([]model.Product, error) {
	var product model.Product
	var products []model.Product
	filter := bson.M{"name": name}
	filterCursor, err := productsCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	var productsFiltered []bson.M
	if err = filterCursor.All(ctx, &productsFiltered); err != nil {
		log.Fatal(err)
	}
	for _, productFiltered := range productsFiltered {
		bsonBytes, _ := bson.Marshal(productFiltered)
		bson.Unmarshal(bsonBytes, &product)
		products = append(products, product)
	}
	return products, err
}

func GetProductByNameLike(name string) ([]model.Product, error) {
	var product model.Product
	var products []model.Product
	filter := bson.M{"name": bson.M{"$regex": "(?i)^"+name}}
	filterCursor, err := productsCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	var productsFiltered []bson.M
	if err = filterCursor.All(ctx, &productsFiltered); err != nil {
		log.Fatal(err)
	}
	for _, productFiltered := range productsFiltered {
		bsonBytes, _ := bson.Marshal(productFiltered)
		bson.Unmarshal(bsonBytes, &product)
		products = append(products, product)
	}
	return products, err
}

func InsertProduct(product *model.Product) {
	product.ID = primitive.NewObjectID()
	insertResult, err := productsCollection.InsertOne(ctx, product)

	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("Inserted post with ID:", insertResult.InsertedID)

}

func UpdateProduct(id string, product *model.Product) {
	fmt.Println(id)
	modelId, _ := primitive.ObjectIDFromHex(id)
	_, err := productsCollection.UpdateOne(
		ctx,
		bson.M{"_id": modelId},
		bson.D{
			{"$set", bson.D{{"name",product.Name},{"price", product.Price}, {"quantity", product.Quantity}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
