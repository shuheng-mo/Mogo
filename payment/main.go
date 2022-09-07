package main

import (
	"github.com/acse-sm321/Mogo/payment/handler"
	log "github.com/micro/go-micro/v2/logger"

	payment "github.com/acse-sm321/Mogo/payment/proto/payment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.payment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), new(handler.Payment))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
