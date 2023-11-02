package dao

import (
	"gorm.io/gorm"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type SetMealDishDao struct {
}

func (s SetMealDishDao) GetBySetMealId(transaction *gorm.DB, SetMealId uint64) ([]model.SetMealDish, error) {
	var dishList []model.SetMealDish
	err := transaction.Where("setmeal_id = ?", SetMealId).Find(&dishList).Error
	return dishList, err
}

func (s SetMealDishDao) InsertBatch(transaction *gorm.DB, setmealDishs []model.SetMealDish) error {
	err := transaction.Create(&setmealDishs).Error
	return err
}

func NewSetMealDishDao() repository.SetMealDishRepo {
	return &SetMealDishDao{}
}
