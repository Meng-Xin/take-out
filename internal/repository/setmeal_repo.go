package repository

import (
	"context"
	"gorm.io/gorm"
	"take-out/internal/model"
)

type SetMealRepo interface {
	Transaction(ctx context.Context) *gorm.DB
	Insert(db *gorm.DB, meal *model.SetMeal) error
}
