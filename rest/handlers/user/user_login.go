package user

import (
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)
type ReqLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	usr, err := h.userRepo.Get(reqLogin.Email, reqLogin.Password)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	if usr == nil {
		http.Error(w, "invalid credentials", http.StatusBadRequest)
		return
	}

	accessToken, err := util.CreateJwt(h.cnf.JwtSecretKey, util.Payload{
		Sub: usr.ID,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		Email: usr.Email,
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	util.SendData(w, accessToken, http.StatusCreated)
}