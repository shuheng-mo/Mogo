package main

import (
	"github.com/acse-sm321/Mogo/cart/handler"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	cart "github.com/acse-sm321/Mogo/cart/proto/cart"
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
