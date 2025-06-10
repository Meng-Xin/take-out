package service

import (
	"context"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"take-out/common"
	"take-out/common/e"
	"take-out/common/enum"
	"take-out/common/retcode"
	"take-out/global"
	"take-out/internal/api/request"
	"take-out/internal/api/response"
	"take-out/internal/model"
	"take-out/internal/repository/dao"
)

type IDishService interface {
	AddDishWithFlavors(ctx context.Context, dto request.DishDTO) error
	PageQuery(ctx context.Context, dto *request.DishPageQueryDTO) (*common.PageResult, error)
	GetByIdWithFlavors(ctx context.Context, id uint64) (response.DishVo, error)
	List(ctx context.Context, categoryId uint64) ([]response.DishListVo, error)
	OnOrClose(ctx context.Context, id uint64, status int) error
	Update(ctx context.Context, dto request.DishUpdateDTO) error
	Delete(ctx context.Context, ids string) error
}

type DishServiceImpl struct {
	repo           *dao.DishDao
	dishFlavorRepo *dao.DishFlavorDao
}

func (d *DishServiceImpl) Delete(ctx context.Context, ids string) error {
	// ids 为多个id的组合，以,进行分割，进行批量删除
	idList := strings.Split(ids, ",")
	for _, idStr := range idList {
		err := global.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			dishId := cast.ToUint64(idStr)
			err := d.dishFlavorRepo.WithTx(tx).DeleteByDishId(ctx, dishId)
			if err != nil {
				global.Log.ErrContext(ctx, "DishServiceImpl.Delete failed, err: %v", err)
				return err
			}
			err = d.repo.WithTx(tx).Delete(ctx, dishId)
			if err != nil {
				global.Log.ErrContext(ctx, "DishServiceImpl.Delete failed, err: %v", err)
				return err
			}
			return nil
		})
		if err != nil {
			global.Log.ErrContext(ctx, "DishServiceImpl.Delete transaction failed, err: %v", err)
			return retcode.NewError(e.MysqlTransActionERR, e.GetMsg(e.MysqlTransActionERR))
		}
	}
	return nil
}

func (d *DishServiceImpl) Update(ctx context.Context, dto request.DishUpdateDTO) error {
	price := cast.ToFloat64(dto.Price)
	dish := model.Dish{
		Id:          dto.Id,
		Name:        dto.Name,
		CategoryId:  dto.CategoryId,
		Price:       price,
		Image:       dto.Image,
		Description: dto.Description,
		Status:      dto.Status,
		Flavors:     dto.Flavors,
	}
	// 开启事务
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 更新菜品信息
		err := d.repo.WithTx(tx).Update(ctx, dish)
		if err != nil {
			global.Log.ErrContext(ctx, "DishServiceImpl.Update failed, err: %v", err)
			return err
		}
		// 更新菜品的口味分两步： 1.先删除原有的所有关联数据，2.再插入新的口味数据
		err = d.dishFlavorRepo.WithTx(tx).DeleteByDishId(ctx, dish.Id)
		if err != nil {
			global.Log.ErrContext(ctx, "DishServiceImpl.Update failed, err: %v", err)
			return err
		}
		if len(dish.Flavors) != 0 {
			err = d.dishFlavorRepo.WithTx(tx).InsertBatch(ctx, dish.Flavors)
			if err != nil {
				global.Log.ErrContext(ctx, "DishServiceImpl.Update failed, err: %v", err)
				return err
			}
		}
		return nil
	})
	if err != nil {
		global.Log.ErrContext(ctx, "DishServiceImpl.Update transaction failed, err: %v", err)
		return retcode.NewError(e.MysqlTransActionERR, e.GetMsg(e.MysqlTransActionERR))
	}
	return err
}

func (d *DishServiceImpl) OnOrClose(ctx context.Context, id uint64, status int) error {
	err := d.repo.OnOrClose(ctx, id, status)
	if err != nil {
		global.Log.ErrContext(ctx, "DishServiceImpl.OnOrClose failed, err: %v", err)
		return err
	}
	return nil
}

func (d *DishServiceImpl) List(ctx context.Context, categoryId uint64) ([]response.DishListVo, error) {
	var dishListVo []response.DishListVo
	dishList, err := d.repo.List(ctx, categoryId)
	if err != nil {
		global.Log.ErrContext(ctx, "DishServiceImpl.List failed, err: %v", err)
		return nil, err
	}
	// 这里这样的写法是 规范化Vo传输对象。
	for _, dish := range dishList {
		dishListVo = append(dishListVo, response.DishListVo{
			Id:          dish.Id,
			Name:        dish.Name,
			CategoryId:  dish.CategoryId,
			Price:       dish.Price,
			Image:       dish.Image,
			Description: dish.Description,
			Status:      dish.Status,
			CreateTime:  dish.CreateTime,
			UpdateTime:  dish.UpdateTime,
			CreateUser:  dish.CreateUser,
			UpdateUser:  dish.UpdateUser,
		})
	}
	return dishListVo, nil
}

func (d *DishServiceImpl) GetByIdWithFlavors(ctx context.Context, id uint64) (response.DishVo, error) {
	dish, err := d.repo.GetById(ctx, id)
	if err != nil {
		global.Log.ErrContext(ctx, "DishServiceImpl.GetByIdWithFlavors failed, err: %v", err)
		return response.DishVo{}, err
	}
	dishVo := response.DishVo{
		Id:          dish.Id,
		Name:        dish.Name,
		CategoryId:  dish.CategoryId,
		Price:       dish.Price,
		Image:       dish.Image,
		Description: dish.Description,
		Status:      dish.Status,
		UpdateTime:  dish.UpdateTime,
		Flavors:     dish.Flavors,
	}
	return dishVo, err
}

func (d *DishServiceImpl) PageQuery(ctx context.Context, dto *request.DishPageQueryDTO) (*common.PageResult, error) {
	pageResult, err := d.repo.PageQuery(ctx, dto)
	if err != nil {
		global.Log.ErrContext(ctx, "DishServiceImpl.PageQuery failed, err: %v", err)
		return nil, err
	}
	return pageResult, err
}

func (d *DishServiceImpl) AddDishWithFlavors(ctx context.Context, dto request.DishDTO) error {
	// 1.构建dish数据
	price, _ := strconv.ParseFloat(dto.Price, 10)
	dish := model.Dish{
		Id:          0,
		Name:        dto.Name,
		CategoryId:  dto.CategoryId,
		Price:       price,
		Image:       dto.Image,
		Description: dto.Description,
		Status:      enum.ENABLE,
	}
	// 开启事务
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 2.先新增菜品数据，再新增口味数据
		err := d.repo.WithTx(tx).Insert(ctx, &dish)
		if err != nil {
			global.Log.ErrContext(ctx, "DishServiceImpl.AddDishWithFlavors failed, err: %v", err)
			return err
		}
		// 为口味数据附加菜品id
		for i, _ := range dto.Flavors {
			dto.Flavors[i].DishId = dish.Id
		}
		err = d.dishFlavorRepo.WithTx(tx).InsertBatch(ctx, dto.Flavors)
		if err != nil {
			global.Log.ErrContext(ctx, "DishServiceImpl.AddDishWithFlavors failed, err: %v", err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.ErrContext(ctx, "DishServiceImpl.AddDishWithFlavors transaction failed, err: %v", err)
		return retcode.NewError(e.MysqlTransActionERR, e.GetMsg(e.MysqlTransActionERR))
	}
	return err
}

func NewDishService(repo *dao.DishDao, dishFlavorRepo *dao.DishFlavorDao) IDishService {
	return &DishServiceImpl{repo: repo, dishFlavorRepo: dishFlavorRepo}
}
