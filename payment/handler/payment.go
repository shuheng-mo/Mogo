package handler

import (
	"github.com/acse-sm321/Mogo/common"
	"github.com/acse-sm321/Mogo/payment/domain/model"
	"github.com/acse-sm321/Mogo/payment/domain/service"
	"context"
	payment "github.com/acse-sm321/Mogo/payment/proto/payment"
)

type Payment struct {
	PaymentDataService service.IPaymentDataService
}

func (e *Payment) AddPayment(ctx context.Context, request *payment.PaymentInfo, response *payment.PaymentID) error{
	payment := &model.Payment{}
	if err := common.SwapTo(request,payment);err!=nil{
		
	}
}

UpdatePayment(ctx context.Context, in *PaymentInfo, out *Response) error
DeletePaymentByID(ctx context.Context, in *PaymentID, out *Response) error
FindPaymentByID(ctx context.Context, in *PaymentID, out *PaymentInfo) error
FindAllPayment(ctx context.Context, in *All, out *PaymentAll) error