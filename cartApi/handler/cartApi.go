package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	cart "github.com/acse-sm321/Mogo/cart/proto/cart"
	cartApi "github.com/acse-sm321/Mogo/cartApi/proto/cartApi"
	log "github.com/micro/go-micro/v2/logger"
	"strconv"
)

type CartApi struct {
	CartService cart.CartService
}

func (e *CartApi) FindAll(ctx context.Context, req *cartApi.Request, rsp *cartApi.Response) error {
	log.Info(" request from /cartApi/findAll ")
	if _, ok := req.Get["user_id"]; !ok {
		//rsp.StatusCode= 500
		return errors.New("invalid arguments")
	}
	userIdString := req.Get["user_id"].Values[0]
	fmt.Println(userIdString)
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	if err != nil {
		fmt.Println(err)
		return err
	}
	cartAll, err := e.CartService.GetAll(context.TODO(), &cart.CartFindAll{UserId: userId})
	b, err := json.Marshal(cartAll)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rsp.StatusCode = 200
	rsp.Body = string(b)
	return nil
}
