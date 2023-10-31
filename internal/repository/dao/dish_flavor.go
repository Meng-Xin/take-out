package dao

import (
	"gorm.io/gorm"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type DishFlavorDao struct {
}

func (d *DishFlavorDao) Update(db *gorm.DB, flavor model.DishFlavor) error {
	err := db.Model(&model.DishFlavor{Id: flavor.Id}).Updates(flavor).Error
	return err
}

func (d *DishFlavorDao) InsertBatch(db *gorm.DB, flavor []model.DishFlavor) error {
	// 批量插入口味数据
	err := db.Create(&flavor).Error
	return err
}

func (d *DishFlavorDao) DeleteByDishId(db *gorm.DB, dishId uint64) error {
	// 根据dishId删除对应的口味数据
	err := db.Where("dish_id = ?", dishId).Delete(&model.DishFlavor{}).Error
	return err
}

func (d *DishFlavorDao) GetByDishId(db *gorm.DB, dishId uint64) ([]model.DishFlavor, error) {
	//TODO implement me
	panic("implement me")
}

// NewDishFlavorDao db 是上个事务创建出来的
func NewDishFlavorDao() repository.DishFlavorRepo {
	return &DishFlavorDao{}
}
