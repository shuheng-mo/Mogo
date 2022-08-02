package handler

import (
	"context"
	common2 "github.com/acse-sm321/Mogo/common"
	"github.com/acse-sm321/Mogo/product/domain/model"
	"github.com/acse-sm321/Mogo/product/domain/service"
	product "github.com/acse-sm321/Mogo/product/proto/product"
)

type Product struct {
	ProductDataService service.IProductDataService
}

//
func (h *Product) AddProduct(ctx context.Context, request *product.ProductInfo, response *product.ResponseProduct) error {
	productAdd := &model.Product{}
	if err := common2.SwapTo(request, productAdd); err != nil {
		return err
	}
	productID, err := h.ProductDataService.AddProduct(productAdd)
	if err != nil {
		return err
	}
	response.ProductId = productID
	return nil
}

//
func (h *Product) FindProductByID(ctx context.Context, request *product.RequestID, response *product.ProductInfo) error {
	productData, err := h.ProductDataService.FindProductByID(request.ProductId)
	if err != nil {
		return err
	}
	if err := common2.SwapTo(productData, response); err != nil {
		return err
	}
	return nil
}

//
func (h *Product) UpdateProduct(ctx context.Context, request *product.ProductInfo, response *product.Response) error {
	productAdd := &model.Product{}
	if err := common2.SwapTo(request, productAdd); err != nil {
		return nil
	}
	err := h.ProductDataService.UpdateProduct(productAdd)
	if err != nil {
		return nil
	}
	response.Msg = "Product updated!"
	return nil
}

//
func (h *Product) DeleteProduct(ctx context.Context, request *product.RequestID, response *product.Response) error {
	if err := h.ProductDataService.DeleteProduct(request.ProductId); err != nil {
		return nil
	}
	response.Msg = "Product deleted!"
	return nil
}

//
func (h *Product) FindAllProduct(ctx context.Context, request *product.RequestAll, response *product.AllProduct) error {
	productAll, err := h.ProductDataService.FindAllProduct()
	if err != nil {
		return nil
	}

	for _, v := range productAll {
		productInfo := &product.ProductInfo{}
		err := common2.SwapTo(v, productInfo)
		if err != nil {
			return err
		}
		response.ProductInfo = append(response.ProductInfo, productInfo)
	}
	return nil
}
