package cmd

import (
	"ecommerce/handlers"
	"ecommerce/util"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))

	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductById))

	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))

	fmt.Println("server running on: 5000")

	globalRouter := util.GlobalRouter(mux)

	err := http.ListenAndServe(":5000", globalRouter)
	if err != nil {
		fmt.Println("error starting the sever", err)
	}
}