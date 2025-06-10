package service

import (
	"context"
	"strconv"
	"take-out/common"
	"take-out/common/enum"
	"take-out/global"
	"take-out/internal/api/request"
	"take-out/internal/model"
	"take-out/internal/repository/dao"
)

type ICategoryService interface {
	AddCategory(ctx context.Context, dto request.CategoryDTO) error
	PageQuery(ctx context.Context, dto request.CategoryPageQueryDTO) (*common.PageResult, error)
	List(ctx context.Context, cate int) ([]model.Category, error)
	DeleteById(ctx context.Context, id uint64) error
	Update(ctx context.Context, dto request.CategoryDTO) error
	SetStatus(ctx context.Context, id uint64, status int) error
}

type CategoryImpl struct {
	repo *dao.CategoryDao
}

func (c *CategoryImpl) SetStatus(ctx context.Context, id uint64, status int) error {
	err := c.repo.SetStatus(ctx, model.Category{
		Id:     id,
		Status: status,
	})
	if err != nil {
		global.Log.ErrContext(ctx, "CategoryImpl.SetStatus error", err)
		return err
	}
	return nil
}

func (c *CategoryImpl) Update(ctx context.Context, dto request.CategoryDTO) error {
	cate, _ := strconv.Atoi(dto.Cate)
	sort, _ := strconv.Atoi(dto.Sort)
	err := c.repo.Update(ctx, model.Category{
		Id:   dto.Id,
		Type: cate,
		Name: dto.Name,
		Sort: sort,
	})
	if err != nil {
		global.Log.ErrContext(ctx, "CategoryImpl.Update error", err)
	}
	return err
}

func (c *CategoryImpl) DeleteById(ctx context.Context, id uint64) error {
	err := c.repo.DeleteById(ctx, id)
	if err != nil {
		global.Log.ErrContext(ctx, "CategoryImpl.DeleteById error", err)
		return err
	}
	return nil
}

func (c *CategoryImpl) List(ctx context.Context, cate int) ([]model.Category, error) {
	list, err := c.repo.List(ctx, cate)
	if err != nil {
		global.Log.ErrContext(ctx, "CategoryImpl.List error", err)
	}
	return list, nil
}

func (c *CategoryImpl) PageQuery(ctx context.Context, dto request.CategoryPageQueryDTO) (*common.PageResult, error) {
	query, err := c.repo.PageQuery(ctx, dto)
	if err != nil {
		global.Log.ErrContext(ctx, "CategoryImpl.PageQuery error", err)
		return nil, err
	}
	return query, nil
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
	if err != nil {
		global.Log.ErrContext(ctx, "CategoryImpl.AddCategory error", err)
		return err
	}
	return nil
}

func NewCategoryService(repo *dao.CategoryDao) ICategoryService {
	return &CategoryImpl{repo: repo}
}
