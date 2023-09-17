package service

import (
	"context"
	"strconv"
	"take-out/common"
	"take-out/common/enum"
	"take-out/internal/api/request"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type ICategoryService interface {
	AddCategory(ctx context.Context, dto request.CategoryDTO) error
	PageQuery(ctx context.Context, dto request.CategoryPageQueryDTO) (*common.PageResult, error)
}

type CategoryImpl struct {
	repo repository.CategoryRepo
}

func (c *CategoryImpl) PageQuery(ctx context.Context, dto request.CategoryPageQueryDTO) (*common.PageResult, error) {
	query, err := c.repo.PageQuery(ctx, dto)
	return query, err
}

func (c *CategoryImpl) AddCategory(ctx context.Context, dto request.CategoryDTO) error {
	// 前端传递的参数是错误的string类型，没办法只能强转了
	cate, _ := strconv.Atoi(dto.Cate)
	sort, _ := strconv.Atoi(dto.Sort)
	// 新增分类
	err := c.repo.Insert(ctx, model.Category{
		Type:   cate,
		Name:   dto.Name,
		Sort:   sort,
		Status: enum.ENABLE,
	})
	return err
}

func NewCategoryService(repo repository.CategoryRepo) ICategoryService {
	return &CategoryImpl{repo: repo}
}
