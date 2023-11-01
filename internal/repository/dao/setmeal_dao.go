package dao

import (
	"context"
	"gorm.io/gorm"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type SetMealDao struct {
	db *gorm.DB
}

func (s *SetMealDao) Transaction(ctx context.Context) *gorm.DB {
	return s.db.WithContext(ctx).Begin()
}

func (s *SetMealDao) Insert(db *gorm.DB, meal model.SetMeal) error {
	err := db.Create(&meal).Error
	return err
}

func NewSetMealDao(db *gorm.DB) repository.SetMealRepo {
	return &SetMealDao{db: db}
}
