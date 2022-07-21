package main

import (
	imooc "cap-imooc/proto/cap"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
)

// CapServer we need a struct to implement the functions in protobuf
type CapServer struct {
}

func (c *CapServer) SayHello(ctx context.Context, req *imooc.SayRequest, res *imooc.SayResponse) error {
	// service logic
	res.Answer = "Crush Go like you crush the fxxking life!"
	return nil
}

func main() {
	// create a new service
	service := micro.NewService(
		micro.Name("cap.imooc.server"),
	)

	// initialization
	service.Init()

	// register the service
	imooc.RegisterCapHandler(service.Server(), new(CapServer))

	// run the service
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
