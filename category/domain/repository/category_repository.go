package repository

import (
	"github.com/acse-sm321/Mogo/category/domain/model"
	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	InitTable() error
	FindCategoryByID(int64) (*model.Category, error)
	CreateCategory(*model.Category) (int64, error)
	DeleteCategoryByID(int64) error
	UpdateCategory(*model.Category) error
	FindAll() ([]model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

//NewCategoryRepository categoryRepository
func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb: db}
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

// InitTable
func (u *CategoryRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Category{}).Error
}

// FindCategoryByID
func (u *CategoryRepository) FindCategoryByID(categoryID int64) (category *model.Category, err error) {
	category = &model.Category{}
	return category, u.mysqlDb.First(category, categoryID).Error
}

// CreateCategory
func (u *CategoryRepository) CreateCategory(category *model.Category) (int64, error) {
	return category.ID, u.mysqlDb.Create(category).Error
}

// DeleteCategoryByID
func (u *CategoryRepository) DeleteCategoryByID(categoryID int64) error {
	return u.mysqlDb.Where("id = ?", categoryID).Delete(&model.Category{}).Error
}

//UpdateCategory
func (u *CategoryRepository) UpdateCategory(category *model.Category) error {
	return u.mysqlDb.Model(category).Update(category).Error
}

// FindAll
func (u *CategoryRepository) FindAll() (categoryAll []model.Category, err error) {
	return categoryAll, u.mysqlDb.Find(&categoryAll).Error
}

// FindCategoryByName
func (u *CategoryRepository) FindCategoryByName(categoryName string) (category *model.Category, err error) {
	category = &model.Category{}
	return category, u.mysqlDb.Where("category_name = ?", categoryName).Find(category).Error
}

func (u *CategoryRepository) FindCategoryByLevel(level uint32) (categorySlice []model.Category, err error) {
	return categorySlice, u.mysqlDb.Where("category_level = ?", level).Find(categorySlice).Error
}

func (u *CategoryRepository) FindCategoryByParent(parent int64) (categorySlice []model.Category, err error) {
	return categorySlice, u.mysqlDb.Where("category_parent = ?", parent).Find(categorySlice).Error
}