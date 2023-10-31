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

func (dd *DishDao) Delete(db *gorm.DB, id uint64) error {
	err := db.Delete(&model.Dish{Id: id}).Error
	return err
}

func (dd *DishDao) Update(db *gorm.DB, dish model.Dish) error {
	err := db.Model(&dish).Updates(dish).Error
	return err
}

func (dd *DishDao) OnOrClose(ctx context.Context, id uint64, status int) error {
	err := dd.db.WithContext(ctx).Model(&model.Dish{Id: id}).Update("status", status).Error
	return err
}

func (dd *DishDao) List(ctx context.Context, categoryId uint64) ([]model.Dish, error) {
	var dishList []model.Dish
	err := dd.db.WithContext(ctx).Where("category_id = ?", categoryId).Find(&dishList).Error
	return dishList, err
}

func (dd *DishDao) GetById(ctx context.Context, id uint64) (*model.Dish, error) {
	var dish model.Dish
	dish.Id = id
	err := dd.db.WithContext(ctx).Preload("Flavors").Find(&dish).Error
	return &dish, err
}

func (dd *DishDao) PageQuery(ctx context.Context, dto *request.DishPageQueryDTO) (*common.PageResult, error) {
	var pageResult common.PageResult
	var dishList []response.DishPageVo
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
		Select("dish.*,c.name as category_name").
		Joins("LEFT OUTER JOIN category c ON c.id = dish.category_id").
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
