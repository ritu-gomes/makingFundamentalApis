package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "invalid path parameter", 400)
		return
	}

	var req ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", 400)
		return
	}

	_, error := h.productRepo.Update(repo.Product{
		Id: id,
		Title: req.Title,
		Description: req.Description,
		Price: req.Price,
		ImageUrl: req.ImageUrl,
	})
	if error != nil {
		http.Error(w, "invalid data", 400)
	}

	util.SendData(w, "successfully updated product", 201)

}