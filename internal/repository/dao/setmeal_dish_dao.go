package dao

import (
	"take-out/global/tx"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type SetMealDishDao struct {
}

func (s SetMealDishDao) GetBySetMealId(transactions tx.Transaction, SetMealId uint64) ([]model.SetMealDish, error) {
	var dishList []model.SetMealDish
	db, err := tx.GetGormDB(transactions)
	if err != nil {
		return nil, err
	}
	err = db.Where("setmeal_id = ?", SetMealId).Find(&dishList).Error
	return dishList, err
}

func (s SetMealDishDao) InsertBatch(transactions tx.Transaction, setmealDishs []model.SetMealDish) error {
	db, err := tx.GetGormDB(transactions)
	if err != nil {
		return err
	}
	err = db.Create(&setmealDishs).Error
	return err
}

func NewSetMealDishDao() repository.SetMealDishRepo {
	return &SetMealDishDao{}
}
