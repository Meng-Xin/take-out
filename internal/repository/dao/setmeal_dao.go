package dao

import (
	"context"
	"gorm.io/gorm"
	"take-out/common"
	"take-out/common/e"
	"take-out/common/retcode"
	"take-out/global"
	"take-out/internal/api/request"
	"take-out/internal/api/response"
	"take-out/internal/model"
)

type SetMealDao struct {
	db *gorm.DB
}

func (s *SetMealDao) DeleteByIds(ctx context.Context, ids ...uint64) error {
	err := s.db.WithContext(ctx).Model(&model.SetMeal{}).Where("id IN ? ", ids).Error
	if err != nil {
		global.Log.ErrContext(ctx, "SetMealDao.DeleteByIds failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "delete setMeal failed")
	}
	return nil
}

func (s *SetMealDao) GetByIdWithDish(ctx context.Context, id uint64) (*model.SetMeal, error) {
	var setMeal model.SetMeal
	err := s.db.WithContext(ctx).First(&setMeal, id).Error
	if err != nil {
		global.Log.ErrContext(ctx, "SetMealDao.GetByIdWithDish failed, err: %v", err)
		return nil, retcode.NewError(e.MysqlERR, "Get setMeal failed")
	}
	return &setMeal, nil
}

func (s *SetMealDao) SetStatus(ctx context.Context, id uint64, status int) error {
	err := s.db.WithContext(ctx).Model(&model.SetMeal{Id: id}).Update("status", status).Error
	if err != nil {
		global.Log.ErrContext(ctx, "SetMealDao.SetStatus failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "Update setMeal failed")
	}
	return nil
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
	err := query.Scopes(pageResult.Paginate(&dto.Page, &dto.PageSize)).
		Select("setmeal.*,c.name as category_name").
		Joins("LEFT JOIN category c ON setmeal.category_id = c.id").
		Order("create_time desc").
		Scan(&setmealPageQueryVo).Error
	if err != nil {
		global.Log.ErrContext(ctx, "SetMealDao.PageQuery List failed, err: %v", err)
		return nil, retcode.NewError(e.MysqlERR, "Get setMeal List failed")
	}
	// 整合数据下发
	pageResult.Records = setmealPageQueryVo
	return &pageResult, nil
}

func (s *SetMealDao) Insert(ctx context.Context, meal *model.SetMeal) error {
	err := s.db.WithContext(ctx).Create(&meal).Error
	if err != nil {
		global.Log.ErrContext(ctx, "SetMealDao.Insert failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "Create setMeal failed")
	}
	return nil
}

func NewSetMealDao(db *gorm.DB) *SetMealDao {
	return &SetMealDao{db: db}
}

func (s *SetMealDao) WithTx(db *gorm.DB) *SetMealDao {
	return &SetMealDao{db: db}
}
