package repository

import (
	"gorm.io/gorm"
	"take-out/internal/model"
)

type DishFlavorRepo interface {
	// InsertBatch 批量插入菜品口味
	InsertBatch(db *gorm.DB, flavor []model.DishFlavor) error
	// DeleteByDishId 根据菜品id删除口味数据
	DeleteByDishId(db *gorm.DB, dishId uint64) error
	// GetByDishId 根据菜品id查询对应的口味数据
	GetByDishId(db *gorm.DB, dishId uint64) ([]model.DishFlavor, error)
	// Update 修改口味数据
	Update(db *gorm.DB, flavor model.DishFlavor) error
}
