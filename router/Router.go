package router

import (
	"github.com/gabrielpsz/dm-challenge/web"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartRouter() {
	fmt.Println("Starting router")
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/products/{name}", web.GetByName).Methods("GET")
	r.HandleFunc("/api/v1/orders", web.Insert).Methods("POST")
	r.HandleFunc("/api/v1/orders", web.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/orders/{id}", web.GetById).Methods("GET")

	var port = ":8090"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}