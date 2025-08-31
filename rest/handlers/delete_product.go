package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "invalid path parameter", 400)
		return
	}

	database.Delete(id)
	util.SendData(w, "successfully deleted the product", 200)
}