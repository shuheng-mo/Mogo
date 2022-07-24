package main

import (
	"github.com/acse-sm321/Mogo/category/handler"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	category "github.com/acse-sm321/Mogo/category/proto/category"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	category.RegisterCategoryHandler(service.Server(), new(handler.Category))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
