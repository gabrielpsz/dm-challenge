package web

import (
	"encoding/json"
	"github.com/gabrielpsz/dm-challenge/model"
	"github.com/gabrielpsz/dm-challenge/repository"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"errors"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var orderRequest model.CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&orderRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var totalOrderPrice float64
	var orderProducts []model.Product
	for _, requestProduct := range orderRequest.Products {
		searchProductList, _ := repository.GetProductByName(requestProduct.Name)
		for _, product := range searchProductList {
			if (product.Quantity != 0) {
				for i := 0; i < requestProduct.Quantity; i++ {
					totalOrderPrice += product.Price
				}
				product.SetOldProductQuantity(product.Quantity)
				product.Quantity = requestProduct.Quantity
				orderProducts = append(orderProducts, product)
			}

		}
	}
	if (len(orderProducts) == 0) {
		message := fmt.Sprintf("The requested products are out of stock")
		e := errors.New(message)
		respondWithError(w, http.StatusInternalServerError, e.Error())
		return
	}
	order := model.NewOrder(orderProducts, totalOrderPrice)

	if err := repository.InsertOrder(order); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, order)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	order := repository.GetOrderById(id)
	if (len(order.Products) == 0) {
		message := fmt.Sprintf("Order with id %v not found", id)
		e := errors.New(message)
		respondWithError(w, http.StatusInternalServerError, e.Error())
		return
	}
	respondWithJson(w, http.StatusOK, order)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	product := repository.GetOrders()
	respondWithJson(w, http.StatusOK, product)
}
