package main

import (
	"git.imooc.com/cap1573/cart/handler"
	log "github.com/micro/go-micro/v2/logger"

	cart "git.imooc.com/cap1573/cart/proto/cart"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	cart.RegisterCartHandler(service.Server(), new(handler.Cart))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
