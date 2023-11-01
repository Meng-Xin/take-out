package dao

import (
	"gorm.io/gorm"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type SetMealDishDao struct {
}

func (s SetMealDishDao) InsertBatch(transaction *gorm.DB, setmealDishs []model.SetMealDish) error {
	err := transaction.Create(&setmealDishs).Error
	return err
}

func NewSetMealDishDao() repository.SetMealDishRepo {
	return &SetMealDishDao{}
}
