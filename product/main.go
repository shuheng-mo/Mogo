package main

import (
	"github.com/acse-sm321/Mogo/product/handler"
	log "github.com/micro/go-micro/v2/logger"

	product "github.com/acse-sm321/Mogo/product/proto/product"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	product.RegisterProductHandler(service.Server(), new(handler.Product))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
