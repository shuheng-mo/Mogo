package main

import (
	"context"
	"fmt"
	go_micro_service_cart "github.com/acse-sm321/Mogo/cart/proto/cart"
	"github.com/acse-sm321/Mogo/cartApi/handler"
	cartApi "github.com/acse-sm321/Mogo/cartApi/proto/cartApi"
	"github.com/acse-sm321/Mogo/common"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"net"
	"net/http"
)

func main() {
	// Consul configuration
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// Jaeger setup
	t, io, err := common.NewTracer("go.micro.api.cartApi", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// hystrix configuration
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()

	// setup hystrix-go
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9096"), hystrixStreamHandler)
		if err != nil {
			log.Error(err)
		}
	}()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.cartApi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8086"),
		micro.Registry(consul),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		// set up hystrix
		micro.WrapClient(NewClientHystrixWrapper()),
		// set up load balance
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	// Initialise service
	service.Init()

	cartService := go_micro_service_cart.NewCartService("go.micro.service.cart", service.Client())

	// Register Handler
	if err := cartApi.RegisterCartApiHandler(service.Server(), &handler.CartApi{CartService: cartService}); err != nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

// Call calling the api exposed
func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		fmt.Println(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(err error) error {
		fmt.Println(err)
		return err
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}
