package main

import (
	"context"
	"fmt"
	common2 "github.com/acse-sm321/Mogo/product/common"
	go_micro_service_product "github.com/acse-sm321/Mogo/product/proto/product"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

func main() {
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// path tracing
	t, io, err := common2.NewTracer("go.micro.service.product.client", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("go.micro.service.product.client"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		micro.Registry(consul),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)

	productService := go_micro_service_product.NewProductService("go.micro.service.product", service.Client())

	productAdd := &go_micro_service_product.ProductInfo{
		ProductName:        "fuck",
		ProductSku:         "bitch",
		ProductPrice:       1.1,
		ProductDescription: "fuck-bitch",
		ProductCategoryId:  1,
		ProductImage: []*go_micro_service_product.ProductImage{
			{
				ImageName: "porno1",
				ImageCode: "pornoimg1",
				ImageUrl:  "pornoimg1",
			},
			{
				ImageName: "porno2",
				ImageCode: "pornoimg2",
				ImageUrl:  "pornoimg2",
			},
		},
		ProductSize: []*go_micro_service_product.ProductSize{
			{
				SizeName: "porno-size",
				SizeCode: "porno-size",
			},
		},
		ProductSeo: &go_micro_service_product.ProductSeo{
			SeoTitle:       "porno-seo",
			SeoKeywords:    "porno-seo",
			SeoDescription: "porno-seo",
			SeoCode:        "porno-seo",
		},
	}
	response, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
