package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")

	id, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "invalid path parameter", 400)
		return
	}

	for _, product := range database.ProductList{
		if product.Id == id {
			util.SendData(w, product, 200)
			return
		}
		util.SendData(w, "id not found", 404)
	}

}