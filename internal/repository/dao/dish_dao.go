package dao

import (
	"context"
	"gorm.io/gorm"
	"take-out/common"
	"take-out/internal/api/request"
	"take-out/internal/api/response"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type DishDao struct {
	db *gorm.DB
}

func (dd *DishDao) PageQuery(ctx context.Context, dto *request.DishPageQueryDTO) (*common.PageResult, error) {
	var pageResult common.PageResult
	var dishList []response.DishVo
	// 1.动态拼接sql
	query := dd.db.WithContext(ctx).Model(&model.Dish{})
	if dto.Name != "" {
		query = query.Where("name LIKE ", "%"+dto.Name+"%")
	}
	if dto.Status != 0 {
		query = query.Where("status = ?", dto.Status)
	}
	if dto.CategoryId != 0 {
		query = query.Where("category_id = ?", dto.CategoryId)
	}
	// 2.动态查询Total
	if err := query.Count(&pageResult.Total).Error; err != nil {
		return nil, err
	}
	// 3.通用分页查询
	if err := query.Scopes(pageResult.Paginate(&dto.Page, &dto.PageSize)).
		Select("*,c.name as category_name").
		Joins("LEFT JOIN category c ON c.id = dish.category_id").
		Order("dish.create_time desc").Scan(&dishList).Error; err != nil {
		return nil, err
	}
	// 构造返回结果
	pageResult.Records = dishList
	return &pageResult, nil
}

// Transaction 开启事务
func (dd *DishDao) Transaction(ctx context.Context) *gorm.DB {
	return dd.db.WithContext(ctx).Begin()
}

// Insert 使用事务指针进行插入菜品数据
func (dd *DishDao) Insert(transaction *gorm.DB, dish *model.Dish) error {
	err := transaction.Create(dish).Error
	return err
}

func NewDishRepo(db *gorm.DB) repository.DishRepo {
	return &DishDao{db: db}
}
