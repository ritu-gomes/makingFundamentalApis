package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migration")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	userService := user.NewService(userRepo)
	productService := product.NewService(productRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := productHandler.NewHandler(middlewares, productService)
	userHandler := userHandler.NewHandler(cnf, userService)

	server := rest.NewServer(
		cnf, 
		productHandler, 
		userHandler)
		
	server.Start()
}