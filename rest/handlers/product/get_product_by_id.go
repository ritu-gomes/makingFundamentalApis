package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "invalid path parameter", 400)
		return
	}

	product, err := h.productRepo.Get(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if product == nil {
		util.SendError(w, 404, "product not found")
		return
	}
	util.SendData(w, product, 200)

}