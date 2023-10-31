package repository

import (
	"context"
	"gorm.io/gorm"
	"take-out/common"
	"take-out/internal/api/request"
	"take-out/internal/model"
)

type DishRepo interface {
	Transaction(ctx context.Context) *gorm.DB
	Insert(db *gorm.DB, dish *model.Dish) error
	PageQuery(ctx context.Context, dto *request.DishPageQueryDTO) (*common.PageResult, error)
	GetById(ctx context.Context, id uint64) (*model.Dish, error)
	List(ctx context.Context, categoryId uint64) ([]model.Dish, error)
	OnOrClose(ctx context.Context, id uint64, status int) error
	Update(db *gorm.DB, dish model.Dish) error
	Delete(db *gorm.DB, id uint64) error
}
