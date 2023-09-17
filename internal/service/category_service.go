package service

import (
	"context"
	"take-out/common/enum"
	"take-out/internal/api/request"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type ICategoryService interface {
	AddCategory(ctx context.Context, dto request.CategoryDTO) error
}

type CategoryImpl struct {
	repo repository.CategoryRepo
}

func (c *CategoryImpl) AddCategory(ctx context.Context, dto request.CategoryDTO) error {
	// 新增分类
	err := c.repo.Insert(ctx, model.Category{
		Cate:   dto.Cate,
		Name:   dto.Name,
		Sort:   dto.Sort,
		Status: enum.ENABLE,
	})
	return err
}

func NewCategoryService(repo repository.CategoryRepo) ICategoryService {
	return &CategoryImpl{repo: repo}
}
