package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid data", 400)
		return
	}
	
	createdProduct, error := h.productRepo.Create(repo.Product{
		Title: req.Title,
		Description: req.Description,
		Price: req.Price,
		ImageUrl: req.ImageUrl,
	})
	if error != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	util.SendData(w, createdProduct, http.StatusCreated)
}