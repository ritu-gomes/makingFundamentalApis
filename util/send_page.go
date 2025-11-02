package util

import "net/http"

type PaginatedData struct {
	Data       any        `json:"data" db:"data"`
	Pagination Pagination `json:"pagination" db:"pagination"`
}
type Pagination struct {
	Limit      int64 `json:"limit" db:"limit"`
	Page       int64 `json:"page" db:"page"`
	TotalItems int64 `json:"totalItems" db:"totalItems"`
	TotalPages int64 `json:"totalPages" db:"totalPages"`
}

func SendPage(w http.ResponseWriter, data any, page, limit, count int64){
	paginatedData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Page: page,
			Limit: limit,
			TotalItems: count,
			TotalPages: count/limit,
		},
	}
	SendData(w, paginatedData, http.StatusOK)
}