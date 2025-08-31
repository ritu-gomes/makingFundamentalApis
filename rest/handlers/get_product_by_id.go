package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "invalid path parameter", 400)
		return
	}

	product := database.Get(id)
	if product == nil {
		util.SendError(w, 404, "product not found")
		return
	}
	util.SendData(w, product, 200)

}