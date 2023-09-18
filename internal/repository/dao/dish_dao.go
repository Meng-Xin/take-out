package dao

import (
	"context"
	"gorm.io/gorm"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type DishDao struct {
	db *gorm.DB
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
