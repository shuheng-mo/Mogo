package handler

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/shuheng-mo/Mogo/category/domain/model"
	"github.com/shuheng-mo/Mogo/category/domain/service"
	category "github.com/shuheng-mo/Mogo/category/proto/category"
	"github.com/shuheng-mo/Mogo/common"
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

// UpdateCategory
func (c *Category) UpdateCategory(ctx context.Context, request *category.CategoryRequest, response *category.UpdateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(response, category)
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(category)
	if err != nil {
		return err
	}
	response.Message = "Category updated!"
	return nil
}

// DeleteCategory
func (c *Category) DeleteCategory(ctx context.Context, request *category.DeleteCategoryRequest, response *category.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(request.CategoryId)
	if err != nil {
		return err
	}
	response.Message = "Category deleted!"
	return nil
}

// FindCategoryByName
func (c *Category) FindCategoryByName(ctx context.Context, request *category.FindByNameRequest, response *category.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)
}

// FindCategoryByID
func (c *Category) FindCategoryByID(ctx context.Context, request *category.FindByIdRequest, response *category.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByID(request.CategoryId)
	if err != nil {
		return err
	}
	return common.SwapTo(category, response)
}

func (c *Category) FindCategoryByLevel(ctx context.Context, request *category.FindByLevelRequest, response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByLevel(request.Level)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, response)
	return nil
}

func (c *Category) FindCategoryByParent(ctx context.Context, request *category.FindByParentRequest, response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByParent(request.ParentId)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, response)
	return nil
}

func (c *Category) FindAllCategory(ctx context.Context, request *category.FindAllRequest, response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, response)
	return nil
}

func categoryToResponse(categorySlice []model.Category, response *category.FindAllResponse) {
	for _, cg := range categorySlice {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(cg, cr)
		if err != nil {
			log.Error(err)
			break
		}
		response.Category = append(response.Category, cr)
	}
}
