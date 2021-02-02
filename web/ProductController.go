package web

import (
	"github.com/gabrielpsz/dm-challenge/repository"
	"github.com/gorilla/mux"
	"net/http"
)

func GetByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	products, err := repository.GetProductByNameLike(name)
	if len(products) == 0 || err != nil {
		respondWithError(w, http.StatusBadRequest, "Product with name" +name+" wasnt found")
		return
	}

	respondWithJson(w, http.StatusOK, products)
}
