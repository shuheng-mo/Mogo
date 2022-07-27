package handler

import (
	"context"
	"git.imooc.com/cap1573/cart/domain/service"
	cart "git.imooc.com/cap1573/cart/proto/cart"
	log "github.com/micro/go-micro/v2/logger"
)

type Cart struct {
	CartDataService service.ICartDataService
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Cart) Call(ctx context.Context, req *cart.Request, rsp *cart.Response) error {
	log.Info("Received Cart.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Cart) Stream(ctx context.Context, req *cart.StreamingRequest, stream cart.Cart_StreamStream) error {
	log.Infof("Received Cart.Stream request with count: %d", req.Count)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&cart.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}
	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Cart) PingPong(ctx context.Context, stream cart.Cart_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&cart.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
