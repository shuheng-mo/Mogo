package main

import (
	"git.imooc.com/cap1573/category/handler"
	log "github.com/micro/go-micro/v2/logger"

	category "git.imooc.com/cap1573/category/proto/category"
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
