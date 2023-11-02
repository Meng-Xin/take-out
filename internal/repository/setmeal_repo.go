package repository

import (
	"context"
	"gorm.io/gorm"
	"take-out/common"
	"take-out/internal/api/request"
	"take-out/internal/model"
)

type SetMealRepo interface {
	Transaction(ctx context.Context) *gorm.DB
	Insert(db *gorm.DB, meal *model.SetMeal) error
	PageQuery(ctx context.Context, dto request.SetMealPageQueryDTO) (*common.PageResult, error)
	SetStatus(ctx context.Context, id uint64, status int) error
	GetByIdWithDish(db *gorm.DB, dishId uint64) (model.SetMeal, error)
}
