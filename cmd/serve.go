package cmd

import (
	"ecommerce/middleware"
	"ecommerce/util"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Hudai)
	mux := http.NewServeMux()

	initRoutes(mux, manager)

	fmt.Println("server running on: 5000")

	globalRouter := util.GlobalRouter(mux)

	err := http.ListenAndServe(":5000", globalRouter)
	if err != nil {
		fmt.Println("error starting the sever", err)
	}
}