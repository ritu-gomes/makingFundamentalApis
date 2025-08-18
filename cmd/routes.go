package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /ritu", 
		manager.With(
			http.HandlerFunc(handlers.Test),
		))

	mux.Handle("GET /products", 
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		))

	mux.Handle("GET /products/{productId}", 
		manager.With(
			http.HandlerFunc(handlers.GetProductById),
		))

	mux.Handle("POST /products", 
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		))
}