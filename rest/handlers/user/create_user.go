package user

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"Email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}
	
	createdUser, err := h.userRepo.Create(repo.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		Password: req.Password,
		IsShopOwner: req.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdUser, http.StatusCreated)
}