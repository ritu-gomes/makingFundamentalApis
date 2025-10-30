package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "invalid path parameter", 400)
		return
	}

	err = h.svc.Delete(id)
	if err != nil {
		http.Error(w, "Internal Server Error", 400)
		return
	}
	util.SendData(w, "successfully deleted the product", 200)
}