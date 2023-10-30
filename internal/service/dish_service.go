package service

import (
	"context"
	"strconv"
	"take-out/common"
	"take-out/common/enum"
	"take-out/internal/api/request"
	"take-out/internal/api/response"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type IDishService interface {
	AddDishWithFlavors(ctx context.Context, dto request.DishDTO) error
	PageQuery(ctx context.Context, dto *request.DishPageQueryDTO) (*common.PageResult, error)
	GetByIdWithFlavors(ctx context.Context, id uint64) (response.DishVo, error)
	List(ctx context.Context, categoryId uint64) ([]response.DishListVo, error)
}

type DishServiceImpl struct {
	repo           repository.DishRepo
	dishFlavorRepo repository.DishFlavorRepo
}

func (d *DishServiceImpl) List(ctx context.Context, categoryId uint64) ([]response.DishListVo, error) {
	var dishListVo []response.DishListVo
	dishList, err := d.repo.List(ctx, categoryId)
	if err != nil {
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
