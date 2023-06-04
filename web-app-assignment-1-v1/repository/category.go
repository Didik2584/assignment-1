package repository

import (
	"a21hc3NpZ25tZW50/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Store(category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (c *categoryRepository) Store(category *model.Category) error {
	err := c.db.Create(category).Error
	return err
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	err := c.db.Model(&model.Category{}).Where("id = ?", id).Updates(category).Error
	return err
}

func (c *categoryRepository) Delete(id int) error {
	err := c.db.Delete(&model.Category{}, id).Error
	return err
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	var category model.Category
	err := c.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	var categories []model.Category
	err := c.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
