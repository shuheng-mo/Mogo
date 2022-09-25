package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	go_micro_service_cart "github.com/shuheng-mo/Mogo/cart/proto/cart"
	"github.com/shuheng-mo/Mogo/cartApi/handler"
	cartApi "github.com/shuheng-mo/Mogo/cartApi/proto/cartApi"
	"github.com/shuheng-mo/Mogo/common"
	"net"
	"net/http"
)

func main() {
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	t, io, err := common.NewTracer("go.micro.api.cartApi", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()

	go func() {
		err = http.ListenAndServe(net.JoinHostPort("146.169.157.226", "9096"), hystrixStreamHandler)
		if err != nil {
			log.Error(err)
		}
	}()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.cartApi"),
		micro.Version("latest"),
		micro.Address("146.169.157.226:8086"),
		//micro.Address("127.0.0.1:8086"),
		micro.Registry(consul),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(NewClientHystrixWrapper()),
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	// Initialise service
	service.Init()

	cartService := go_micro_service_cart.NewCartService("go.micro.service.cart", service.Client())

	//cartService.AddCart(context.TODO(), &go_micro_service_cart.CartInfo{
	//	UserId:    3,
	//	ProductId: 4,
	//	SizeId:    5,
	//	Num:       5,
	//})

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

// to do the registration
// docker run --rm -p 8080:8080 -e ICODE=68272CAA3468CEED cap1573/cap-micro --registry=consul --registry_address=YOUR_IP:8500 api --handler=api
// run the hystrix dashboard
// docker run -d -p 9002:9002 cap1573/hystrix-dashboard
// access the hystrix dashboard by http://localhost:9002/hystrix
// access the hystrix stream: http://146.169.157.226:9096/hystrix.stream
// access the api by: http://127.0.0.1:8080/cartApi/findAll?user_id=1
//146.169.157.226.
