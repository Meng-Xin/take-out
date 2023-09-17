package dao

import (
	"context"
	"gorm.io/gorm"
	"take-out/common"
	"take-out/internal/api/request"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type CategoryDao struct {
	db *gorm.DB
}

func (c *CategoryDao) PageQuery(ctx context.Context, dto request.CategoryPageQueryDTO) (*common.PageResult, error) {
	var pageResult common.PageResult
	var categoryList []model.Category

	// 构造初始查询结构
	query := c.db.WithContext(ctx).Model(&model.Category{})
	if dto.Name != "" {
		query = query.Where("name like ?", "%"+dto.Name+"%")
	}
	if dto.Cate != 0 {
		query = query.Where("type = ?", dto.Cate)
	}
	// 计算总数
	if err := query.Count(&pageResult.Total).Error; err != nil {
		return nil, err
	}
	// 查询数据
	err := query.Scopes(pageResult.Paginate(&dto.Page, &dto.PageSize)).Find(&categoryList).Error
	pageResult.Records = categoryList
	return &pageResult, err
}

func (c *CategoryDao) Insert(ctx context.Context, category model.Category) error {
	// 新增分类数据
	err := c.db.WithContext(ctx).Create(&category).Error
	return err
}

func NewCategoryDao(db *gorm.DB) repository.CategoryRepo {
	return &CategoryDao{db: db}
}
