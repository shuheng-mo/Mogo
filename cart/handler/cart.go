package handler

import (
	"context"
	"github.com/acse-sm321/Mogo/cart/domain/model"
	"github.com/acse-sm321/Mogo/cart/domain/service"
	cart "github.com/acse-sm321/Mogo/cart/proto/cart"
	"github.com/acse-sm321/Mogo/common"
)

type Cart struct {
	CartDataService service.ICartDataService
}

// AddCart add new cart
func (h *Cart) AddCart(ctx context.Context, request *cart.CartInfo, response *cart.ResponseAdd) (err error) {
	cart := &model.Cart{}
	common.SwapTo(request, cart)
	response.CartId, err = h.CartDataService.AddCart(cart)
	return err
}

func (h *Cart) CleanCart(ctx context.Context, request *cart.Clean, response *cart.Response) (err error) {
	if err := h.CartDataService.CleanCart(request.UserId); err != nil {
		return err
	}
	response.Meg = "Cart cleaned successfully"
	return nil
}

func (h *Cart) Incr(ctx context.Context, request *cart.Item, response *cart.Response) (err error) {
	if err := h.CartDataService.IncrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Meg = "Increased the num of the cart"
	return nil
}

func (h *Cart) Decr(ctx context.Context, request *cart.Item, response *cart.Response) (err error) {
	if err := h.CartDataService.DecrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Meg = "Decreased the num of the cart"
	return nil
}

func (h *Cart) DeleteItemByID(ctx context.Context, request *cart.CartID, response *cart.Response) (err error) {
	if err := h.CartDataService.DeleteCart(request.Id); err != nil {
		return err
	}
	response.Meg = "Cart has been deleted"
	return nil
}

// GetAll get all carts of a user
func (h *Cart) GetAll(ctx context.Context, request *cart.CartFindAll, response *cart.CartAll) (err error) {
	cartAll, err := h.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}

	for _, v := range cartAll {
		cart := &cart.CartInfo{}
		if err := common.SwapTo(v, cart); err != nil {
			return err
		}
		response.CartInfo = append(response.CartInfo, cart)
	}
	return nil
}
