package dao

import (
	"context"
	"gorm.io/gorm"
	"take-out/common/e"
	"take-out/common/retcode"
	"take-out/global"
	"take-out/internal/model"
)

type DishFlavorDao struct {
	db *gorm.DB
}

func (d *DishFlavorDao) Update(ctx context.Context, flavor model.DishFlavor) error {
	err := d.db.WithContext(ctx).Model(&model.DishFlavor{Id: flavor.Id}).Updates(flavor).Error
	if err != nil {
		global.Log.ErrContext(ctx, "Update dish flavor failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "update dish flavor failed")
	}
	return nil
}

func (d *DishFlavorDao) InsertBatch(ctx context.Context, flavor []model.DishFlavor) error {
	// 批量插入口味数据
	err := d.db.WithContext(ctx).Create(&flavor).Error
	if err != nil {
		global.Log.ErrContext(ctx, "DishFlavorDao.InsertBatch dish flavor failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "insert dish flavor failed")
	}
	return nil
}

func (d *DishFlavorDao) DeleteByDishId(ctx context.Context, dishId uint64) error {
	// 根据dishId删除对应的口味数据
	err := d.db.WithContext(ctx).Where("dish_id = ?", dishId).Delete(&model.DishFlavor{}).Error
	if err != nil {
		global.Log.ErrContext(ctx, "DishFlavorDao.DeleteByDishId dish flavor failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "delete dish flavor failed")
	}
	return nil
}

func (d *DishFlavorDao) GetByDishId(ctx context.Context, dishId uint64) ([]model.DishFlavor, error) {
	//TODO implement me
	panic("implement me")
}

// NewDishFlavorDao 创建dao实例
func NewDishFlavorDao(db *gorm.DB) *DishFlavorDao {
	return &DishFlavorDao{db: db}
}

// WithTx 使用事务
func (d *DishFlavorDao) WithTx(tx *gorm.DB) *DishFlavorDao {
	return &DishFlavorDao{db: tx}
}
