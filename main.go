package main

import (
	"ecommerce/util"
	"fmt"
)

func main() {
	// cnf := config.GetConfig()

	// fmt.Println(cnf.Version)
	// fmt.Println(cnf.ServiceName)
	// fmt.Println(cnf.HttpPort)

	// cmd.Serve()
	jwt, err := util.CreateJwt("my_secret", util.Payload{
		Sub: 45,
		FirstName: "Irene",
		LastName: "Ritu",
		Email: "ritzy333@gmail.com",
		IsShopOwner: false,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(jwt)

}