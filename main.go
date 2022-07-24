package main

import (
	"fmt"
	"github.com/acse-sm321/Mogo/domain/repository"
	service2 "github.com/acse-sm321/Mogo/domain/service"
	user "github.com/acse-sm321/Mogo/proto/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"

	"github.com/acse-sm321/Mogo/handler"
)

// get the
// docker run --rm -v $(pwd):$(pwd) -w $(pwd) -e ICODE={CODE_HERE} cap1573/cap-protoc -I ./ --go_out=./ --micro_out=./ ./proto/user/user.proto

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	// initialization
	srv.Init()

	// create MySQL
	db, err := gorm.Open("mysql", "root:19990429@/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	// close database after use
	defer db.Close()

	db.SingularTable(true)

	// create repository (Init Table for only once at the first time)
	//rp := repository.NewUserRepository(db)
	//rp.InitTable()

	// create the service instance
	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))
	// register the handler
	err = user.RegisterUserHandler(srv.Server(), &handler.User{userDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Register handler
	//micro.Handle(new(handler.User))

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
