package dao

import (
	"gorm.io/gorm"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type DishFlavorDao struct {
}

func (d *DishFlavorDao) InsertBatch(db *gorm.DB, flavor []model.DishFlavor) error {
	// 批量插入口味数据
	err := db.Create(&flavor).Error
	return err
}

func (d *DishFlavorDao) DeleteByDishId(db *gorm.DB, dishId uint64) error {
	//TODO implement me
	panic("implement me")
}

func (d *DishFlavorDao) GetByDishId(db *gorm.DB, dishId uint64) ([]model.DishFlavor, error) {
	//TODO implement me
	panic("implement me")
}

// NewDishFlavorDao db 是上个事务创建出来的
func NewDishFlavorDao() repository.DishFlavorRepo {
	return &DishFlavorDao{}
}
