package dao

import (
	"context"
	"gorm.io/gorm"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type CategoryDao struct {
	db *gorm.DB
}

func (c *CategoryDao) Insert(ctx context.Context, category model.Category) error {
	// 新增分类数据
	err := c.db.WithContext(ctx).Create(&category).Error
	return err
}

func NewCategoryDao(db *gorm.DB) repository.CategoryRepo {
	return &CategoryDao{db: db}
}
