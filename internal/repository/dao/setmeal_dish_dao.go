package dao

import (
	"context"
	"gorm.io/gorm"
	"take-out/common/e"
	"take-out/common/retcode"
	"take-out/internal/model"
)

type SetMealDishDao struct {
	db *gorm.DB
}

func (d *SetMealDishDao) GetBySetMealId(ctx context.Context, SetMealId uint64) ([]model.SetMealDish, error) {
	var dishList []model.SetMealDish
	err := d.db.WithContext(ctx).Where("setmeal_id = ?", SetMealId).Find(&dishList).Error
	if err != nil {
		return nil, retcode.NewError(e.MysqlERR, "delete dish failed")
	}
	return dishList, nil
}

func (d *SetMealDishDao) InsertBatch(ctx context.Context, setmealDishs []model.SetMealDish) error {
	err := d.db.WithContext(ctx).Create(&setmealDishs).Error
	if err != nil {
		return retcode.NewError(e.MysqlERR, "delete dish failed")
	}
	return nil
}

func NewSetMealDishDao() *SetMealDishDao {
	return &SetMealDishDao{}
}

func (d *SetMealDishDao) WithTx(db *gorm.DB) *SetMealDishDao {
	return &SetMealDishDao{db: db}
}
