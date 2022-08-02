package main

import (
	"github.com/acse-sm321/Mogo/category/domain/repository"
	service2 "github.com/acse-sm321/Mogo/category/domain/service"
	"github.com/acse-sm321/Mogo/category/handler"
	category "github.com/acse-sm321/Mogo/category/proto/category"
	"github.com/acse-sm321/Mogo/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// Configuration of consul
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// Register consul
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		// set IP address and ports to expose
		micro.Address("127.0.0.1:8082"),
		// Add consul as registry center
		micro.Registry(consulRegistry),
	)

	// get mysql config (no prefix needed)
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")

	// Connect database
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()

	// No duplicated table
	db.SingularTable(true)

	// runs for only once
	// rp := repository.NewCategoryRepository(db)
	// rp.InitTable()

	// Initialise service
	service.Init()

	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))

	// Register Handler
	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{CategoryDataService: categoryDataService})
	if err != nil {
		log.Error(err)
	}

	//category.RegisterCategoryHandler(service.Server(), new(handler.Category))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
