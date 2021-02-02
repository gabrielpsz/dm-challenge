package web

import (
	"github.com/gabrielpsz/dm-challenge/repository"
	"github.com/gorilla/mux"
	"net/http"
)

func GetByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product, err := repository.GetProductByName(params["name"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product Name")
		return
	}
	respondWithJson(w, http.StatusOK, product)
}
