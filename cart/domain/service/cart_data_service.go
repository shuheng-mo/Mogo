package service

import (
	"github.com/acse-sm321/Mogo/cart/domain/model"
	"github.com/acse-sm321/Mogo/cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	DecrNum(int64, int64) error
	IncrNum(int64, int64) error
}

func NewCartDataService(cartRepository repository.ICartRepository) ICartDataService {
	return &CartDataService{cartRepository}
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

func (u *CartDataService) AddCart(cart *model.Cart) (int64, error) {
	return u.CartRepository.CreateCart(cart)
}

func (u *CartDataService) DeleteCart(cartID int64) error {
	return u.CartRepository.DeleteCartByID(cartID)
}

func (u *CartDataService) UpdateCart(cart *model.Cart) error {
	return u.CartRepository.UpdateCart(cart)
}

func (u *CartDataService) FindCartByID(cartID int64) (*model.Cart, error) {
	return u.CartRepository.FindCartByID(cartID)
}

func (u *CartDataService) FindAllCart(userID int64) ([]model.Cart, error) {
	return u.CartRepository.FindAll(userID)
}

func (u *CartDataService) CleanCart(userID int64) error {
	return u.CartRepository.CleanCart(userID)
}

func (u *CartDataService) DecrNum(cartID int64, num int64) error {
	return u.CartRepository.DecrNum(cartID, num)
}

func (u *CartDataService) IncrNum(cartID int64, num int64) error {
	return u.CartRepository.IncrNum(cartID, num)
}
