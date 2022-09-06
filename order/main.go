package main

import (
	"github.com/acse-sm321/Mogo/common"
	repository2 "github.com/acse-sm321/Mogo/order/domain/repository"
	service2 "github.com/acse-sm321/Mogo/order/domain/service"
	"github.com/acse-sm321/Mogo/order/handler"
	order "github.com/acse-sm321/Mogo/order/proto/order"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/opentracing/opentracing-go"
)

var (
	QPS = 1000 // max query per sesond
)

func main() {
	// Consul
	consulCofig, err := common.GetConsulConfig("localhost", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// Register center
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"localhost"}
	})

	// Jaeger
	t, io, err := common.NewTracer("go.micro.service.order", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// Init DB
	mysqlInfo := common.GetMysqlFromConsul(consulCofig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	db.SingularTable(true) // no duplicates

	tableInit := repository2.NewOrderRepository(db)
	tableInit.InitTable()

	// service instance
	orderDataService := service2.NewOrderDataService(repository2.NewOrderRepository(db))

	// expose
	//common.Prometheus

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.order"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	order.RegisterOrderHandler(service.Server(), new(handler.Order))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
