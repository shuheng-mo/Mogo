package handler

import (
	"context"
	"github.com/acse-sm321/Mogo/category/common"
	"github.com/acse-sm321/Mogo/category/domain/model"
	"github.com/acse-sm321/Mogo/category/domain/service"
	category "github.com/acse-sm321/Mogo/category/proto/category"
)

type Category struct {
	CategoryDataService service.ICategoryDataService
}

// CreateCategory provide service for adding new category
func (c *Category) CreateCategory(ctx context.Context, request *category.CategoryRequest, response *category.CreateCategoryResponse) error {
	category := &model.Category{}
	// assignment
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}
	categoryId, err := c.CategoryDataService.AddCategory(category)
	if err != nil {
		return err
	}
	response.Message = "New category added!"
	response.CategoryId = categoryId
	return nil
}
