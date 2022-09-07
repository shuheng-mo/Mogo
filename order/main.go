package main

import (
	"github.com/acse-sm321/Mogo/common"
	"github.com/acse-sm321/Mogo/order/domain/repository"
	service2 "github.com/acse-sm321/Mogo/order/domain/service"
	"github.com/acse-sm321/Mogo/order/handler"
	order "github.com/acse-sm321/Mogo/order/proto/order"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

var (
	QPS = 1000 // maximum query per second
)

func main() {
	consulConfig, err := common.GetConsulConfig("localhost", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"localhost:8500",
		}
	})

	t, io, err := common.NewTracer("go.micro.service.order", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	//tableInit := repository.NewOrderRepository(db)
	//tableInit.InitTable()

	orderDataService := service2.NewOrderDataService(repository.NewOrderRepository(db))

	common.PrometheusBoot(9092)

	service := micro.NewService(
		micro.Name("go.micro.service.order"),
		micro.Version("latest"),
		micro.Address(":9085"),
		micro.Registry(consul),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()

	// Register Handler
	order.RegisterOrderHandler(service.Server(), &handler.Order{OrderDataService: orderDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
