package service

import (
	"context"
	"strconv"
	"strings"
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
	OnOrClose(ctx context.Context, id uint64, status int) error
	Update(ctx context.Context, dto request.DishUpdateDTO) error
	Delete(ctx context.Context, ids string) error
}

type DishServiceImpl struct {
	repo           repository.DishRepo
	dishFlavorRepo repository.DishFlavorRepo
}

func (d *DishServiceImpl) Delete(ctx context.Context, ids string) error {
	// ids 为多个id的组合，以,进行分割，进行批量删除
	idList := strings.Split(ids, ",")
	for _, idStr := range idList {
		// 这里因为循环内部的事务提交，使用匿名函数解决内存泄漏问题。
		err := func() error {
			dishId, _ := strconv.ParseUint(idStr, 10, 64)
			// 开启事务
			transaction := d.repo.Transaction(ctx)
			defer func() {
				if r := recover(); r != nil {
					transaction.Rollback()
				}
			}()
			// 关联删除菜品口味数据
			err := d.dishFlavorRepo.DeleteByDishId(transaction, dishId)
			if err != nil {
				return err
			}
			// 删除菜品
			err = d.repo.Delete(transaction, dishId)
			if err != nil {
				return err
			}
			return transaction.Commit().Error
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DishServiceImpl) Update(ctx context.Context, dto request.DishUpdateDTO) error {
	price, _ := strconv.ParseFloat(dto.Price, 10)
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
	transaction := d.repo.Transaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			transaction.Rollback()
		}
	}()
	// 更新菜品信息
	err := d.repo.Update(transaction, dish)
	if err != nil {
		return err
	}
	// 更新菜品的口味分两步： 1.先删除原有的所有关联数据，2.再插入新的口味数据
	err = d.dishFlavorRepo.DeleteByDishId(transaction, dish.Id)
	if err != nil {
		return err
	}
	if len(dish.Flavors) != 0 {
		err = d.dishFlavorRepo.InsertBatch(transaction, dish.Flavors)
		if err != nil {
			return err
		}
	}

	return transaction.Commit().Error
}

func (d *DishServiceImpl) OnOrClose(ctx context.Context, id uint64, status int) error {
	err := d.repo.OnOrClose(ctx, id, status)
	return err
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
