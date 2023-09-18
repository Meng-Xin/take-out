package repository

import (
	"context"
	"gorm.io/gorm"
	"take-out/internal/model"
)

type DishRepo interface {
	Transaction(ctx context.Context) *gorm.DB
	Insert(db *gorm.DB, dish *model.Dish) error
}
