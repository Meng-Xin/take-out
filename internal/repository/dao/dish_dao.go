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

type DishDao struct {
	db *gorm.DB
}

func (d *DishDao) Delete(ctx context.Context, id uint64) error {
	err := d.db.Delete(&model.Dish{Id: id}).Error
	if err != nil {
		global.Log.ErrContext(ctx, "DishDao.Delete failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "delete dish failed")
	}
	return nil
}

func (d *DishDao) Update(ctx context.Context, dish model.Dish) error {
	err := d.db.Model(&dish).Updates(dish).Error
	if err != nil {
		global.Log.ErrContext(ctx, "DishDao.Update failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "update dish failed")
	}
	return nil
}

func (d *DishDao) OnOrClose(ctx context.Context, id uint64, status int) error {
	err := d.db.WithContext(ctx).Model(&model.Dish{Id: id}).Update("status", status).Error
	if err != nil {
		global.Log.ErrContext(ctx, "DishDao.OnOrClose failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "update dish failed")
	}
	return nil
}

func (d *DishDao) List(ctx context.Context, categoryId uint64) ([]model.Dish, error) {
	var dishList []model.Dish
	err := d.db.WithContext(ctx).Where("category_id = ?", categoryId).Find(&dishList).Error
	if err != nil {
		global.Log.ErrContext(ctx, "DishDao.List failed, err: %v", err)
		return nil, retcode.NewError(e.MysqlERR, "get dish list failed")
	}
	return dishList, nil
}

func (d *DishDao) GetById(ctx context.Context, id uint64) (*model.Dish, error) {
	var dish model.Dish
	dish.Id = id
	err := d.db.WithContext(ctx).Preload("Flavors").Find(&dish).Error
	if err != nil {
		global.Log.ErrContext(ctx, "DishDao.GetById failed, err: %v", err)
		return nil, retcode.NewError(e.MysqlERR, "get dish failed")
	}
	return &dish, nil
}

func (d *DishDao) PageQuery(ctx context.Context, dto *request.DishPageQueryDTO) (*common.PageResult, error) {
	var pageResult common.PageResult
	var dishList []response.DishPageVo
	// 1.动态拼接sql
	query := d.db.WithContext(ctx).Model(&model.Dish{})
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
		global.Log.ErrContext(ctx, "DishDao.PageQuery failed, err: %v", err)
		return nil, retcode.NewError(e.MysqlERR, "get dish failed")
	}
	// 3.通用分页查询
	err := query.Scopes(pageResult.Paginate(&dto.Page, &dto.PageSize)).
		Select("dish.*,c.name as category_name").
		Joins("LEFT OUTER JOIN category c ON c.id = dish.category_id").
		Order("dish.create_time desc").Scan(&dishList).Error
	if err != nil {
		global.Log.ErrContext(ctx, "DishDao.PageQuery failed, err: %v", err)
		return nil, retcode.NewError(e.MysqlERR, "get dish failed")
	}
	// 构造返回结果
	pageResult.Records = dishList
	return &pageResult, nil
}

// Insert 插入菜品数据
func (d *DishDao) Insert(ctx context.Context, dish *model.Dish) error {
	err := d.db.Create(dish).Error
	if err != nil {
		global.Log.ErrContext(ctx, "DishDao.Insert failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "get dish failed")
	}
	return nil
}

func NewDishRepo(db *gorm.DB) *DishDao {
	return &DishDao{db: db}
}

// WithTx 切换到事务模式
func (d *DishDao) WithTx(db *gorm.DB) *DishDao {
	return &DishDao{db: db}
}
