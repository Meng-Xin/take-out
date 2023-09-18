package service

import (
	"context"
	"strconv"
	"take-out/common/enum"
	"take-out/internal/api/request"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type IDishService interface {
	AddDishWithFlavors(ctx context.Context, dto request.DishDTO) error
}

type DishServiceImpl struct {
	repo           repository.DishRepo
	dishFlavorRepo repository.DishFlavorRepo
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
	// 开启事务，出现问题直接回滚
	transaction := d.repo.Transaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()
	// 2.先新增菜品数据，再新增口味数据
	if err := d.repo.Insert(transaction, &dish); err != nil {
		return err
	}
	// 为口味数据附加菜品id
	for i, _ := range dto.Flavors {
		dto.Flavors[i].DishId = dish.Id
	}
	// 3.使用返回的事务指针动态构建 dishFlavor，因为它只依附于动态返回的事务指针。
	if err := d.dishFlavorRepo.InsertBatch(transaction, dto.Flavors); err != nil {
		return err
	}
	return transaction.Commit().Error
}

func NewDishService(repo repository.DishRepo, dishFlavorRepo repository.DishFlavorRepo) IDishService {
	return &DishServiceImpl{repo: repo, dishFlavorRepo: dishFlavorRepo}
}
