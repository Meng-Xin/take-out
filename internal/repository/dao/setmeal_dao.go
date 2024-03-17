package dao

import (
	"context"
	"gorm.io/gorm"
	"take-out/common"
	"take-out/global/tx"
	"take-out/internal/api/request"
	"take-out/internal/api/response"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type SetMealDao struct {
	db *gorm.DB
}

func (s *SetMealDao) GetByIdWithDish(transactions tx.Transaction, id uint64) (model.SetMeal, error) {
	db, err := tx.GetGormDB(transactions)
	if err != nil {
		return model.SetMeal{}, err
	}
	var setMeal model.SetMeal
	err = db.First(&setMeal, id).Error
	return setMeal, err
}

func (s *SetMealDao) SetStatus(ctx context.Context, id uint64, status int) error {
	err := s.db.WithContext(ctx).Model(&model.SetMeal{Id: id}).Update("status", status).Error
	return err
}

func (s *SetMealDao) PageQuery(ctx context.Context, dto request.SetMealPageQueryDTO) (*common.PageResult, error) {
	var pageResult common.PageResult
	var setmealPageQueryVo []response.SetMealPageQueryVo
	// 构造基础查询
	query := s.db.WithContext(ctx).Model(&model.SetMeal{})
	// 动态构造查询条件
	if dto.CategoryId != 0 {
		query = query.Where("setmeal.category_id = ?", dto.CategoryId)
	}
	if dto.Name != "" {
		query = query.Where("setmeal.name LIKE ?", "%"+dto.Name+"%")
	}
	if dto.Status != 0 {
		query = query.Where("setmeal.status = ?", dto.Status)
	}
	if err := query.Count(&pageResult.Total).Error; err != nil {
		return nil, err
	}
	// 分页查询构造
	if err := query.Scopes(pageResult.Paginate(&dto.Page, &dto.PageSize)).
		Select("setmeal.*,c.name as category_name").
		Joins("LEFT JOIN category c ON setmeal.category_id = c.id").
		Order("create_time desc").
		Scan(&setmealPageQueryVo).Error; err != nil {
		return nil, err
	}
	// 整合数据下发
	pageResult.Records = setmealPageQueryVo
	return &pageResult, nil
}

func (s *SetMealDao) Transaction(ctx context.Context) tx.Transaction {
	return tx.NewGormTransaction(s.db, ctx)
}

func (s *SetMealDao) Insert(transactions tx.Transaction, meal *model.SetMeal) error {
	db, err := tx.GetGormDB(transactions)
	if err != nil {
		return err
	}
	err = db.Create(&meal).Error
	return err
}

func NewSetMealDao(db *gorm.DB) repository.SetMealRepo {
	return &SetMealDao{db: db}
}
