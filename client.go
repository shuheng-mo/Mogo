package main

import (
	imooc "cap-imooc/proto/cap"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
)

func main() {
	// instance
	service := micro.NewService(
		micro.Name("cap.imooc.client"),
	)

	// init
	service.Init()

	capImooc := imooc.NewCapService("cap.imooc.server", service.Client())

	_, err := capImooc.SayHello(context.TODO(), &imooc.SayRequest{
		Message: "Yes, Yes, fuck you too !",
	})

	if err != nil {
		fmt.Println(err)
	}
}
