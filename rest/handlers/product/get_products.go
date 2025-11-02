package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	reqQuery := r.URL.Query()

	pageAsString := reqQuery.Get("page")
	limitAsString := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsString, 10, 32)
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.ParseInt(limitAsString, 10, 32)
	if limit == 0 {
		limit = 10
	}
	
	productList, err := h.svc.List(page, limit)
	if err != nil {
		http.Error(w, "somossa", 400)
		return
	}

	count, err := h.svc.Count()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "internal server error")
	}

	util.SendPage(w, productList, page, limit, count)
}