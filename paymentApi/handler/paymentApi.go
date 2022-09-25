package handler

import (
	"context"
	"encoding/json"
	log "github.com/micro/go-micro/v2/logger"

	paymentApi "github.com/shuheng-mo/Mogo/paymentApi/proto/paymentApi"
)

type PaymentApi struct{}

func (e *PaymentApi) Call(ctx context.Context, req *paymentApi.Request, rsp *paymentApi.Response) error {
	log.Info("Received PaymentApi.Call request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{}")
	rsp.Body = string(b)
	return nil
}
