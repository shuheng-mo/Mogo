package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"

	"github.com/acse-sm321/Mogo/handler"
)

// get the
// docker run --rm -v $(pwd):$(pwd) -w $(pwd) -e ICODE={CODE_HERE} cap1573/cap-protoc -I ./ --go_out=./ --micro_out=./ ./proto/user/user.proto

func main() {
	// Create service
	srv := service.New(
		service.Name("src"),
	)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
