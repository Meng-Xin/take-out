package dao

import (
	"context"
	"gorm.io/gorm"
	"take-out/common"
	"take-out/common/e"
	"take-out/common/retcode"
	"take-out/global"
	"take-out/internal/api/request"
	"take-out/internal/model"
)

type CategoryDao struct {
	db *gorm.DB
}

func (d *CategoryDao) SetStatus(ctx context.Context, category model.Category) error {
	err := d.db.WithContext(ctx).Model(&category).Update("status", category.Status).Error
	if err != nil {
		global.Log.ErrContext(ctx, "SetStatus failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "SetStatus category failed")
	}
	return nil
}

func (d *CategoryDao) Update(ctx context.Context, category model.Category) error {
	err := d.db.WithContext(ctx).Model(&category).Updates(&category).Error
	if err != nil {
		global.Log.ErrContext(ctx, "Update failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "Update category failed")
	}
	return nil
}

func (d *CategoryDao) DeleteById(ctx context.Context, id uint64) error {
	err := d.db.WithContext(ctx).Delete(&model.Category{Id: id}).Error
	if err != nil {
		global.Log.ErrContext(ctx, "DeleteById failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "Delete category failed")
	}
	return nil
}

func (d *CategoryDao) List(ctx context.Context, cate int) ([]model.Category, error) {
	var categoryList []model.Category
	err := d.db.WithContext(ctx).Where("type = ?", cate).
		Order("sort asc").
		Order("create_time desc").
		Find(&categoryList).
		Error
	if err != nil {
		global.Log.ErrContext(ctx, "List failed, err: %v", err)
		return nil, retcode.NewError(e.MysqlERR, "Get category List failed")
	}
	return categoryList, nil
}

func (d *CategoryDao) PageQuery(ctx context.Context, dto request.CategoryPageQueryDTO) (*common.PageResult, error) {
	var pageResult common.PageResult
	var categoryList []model.Category

	// 构造初始查询结构
	query := d.db.WithContext(ctx).Model(&model.Category{})
	if dto.Name != "" {
		query = query.Where("name like ?", "%"+dto.Name+"%")
	}
	if dto.Cate != 0 {
		query = query.Where("type = ?", dto.Cate)
	}
	// 计算总数
	if err := query.Count(&pageResult.Total).Error; err != nil {
		global.Log.ErrContext(ctx, "PageQuery List failed, err: %v", err)
		return nil, retcode.NewError(e.MysqlERR, "Get category List failed")
	}
	// 查询数据
	err := query.Scopes(pageResult.Paginate(&dto.Page, &dto.PageSize)).
		Order("create_time desc").
		Find(&categoryList).
		Error
	if err != nil {
		global.Log.ErrContext(ctx, "PageQuery List failed, err: %v", err)
		return nil, retcode.NewError(e.MysqlERR, "Get category List failed")
	}
	pageResult.Records = categoryList
	return &pageResult, nil
}

func (d *CategoryDao) Insert(ctx context.Context, category model.Category) error {
	// 新增分类数据
	err := d.db.WithContext(ctx).Create(&category).Error
	if err != nil {
		global.Log.ErrContext(ctx, "Insert failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "Create category failed")
	}
	return err
}

func NewCategoryDao(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db: db}
}
