package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImageUrl    string  `json:"image_url" db:"image_url"`
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
		// util.SendError(w, http.StatusInternalServerError, )
		fmt.Println(error)
		http.Error(w, "internal server error create_product_func", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdProduct, http.StatusCreated)
}