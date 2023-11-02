package service

import (
	"context"
	"strconv"
	"take-out/common"
	"take-out/common/enum"
	"take-out/internal/api/request"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type ISetMealService interface {
	SaveWithDish(ctx context.Context, dto request.SetMealDTO) error
	PageQuery(ctx context.Context, dto request.SetMealPageQueryDTO) (*common.PageResult, error)
	OnOrClose(ctx context.Context, id uint64, status int) error
}

type SetMealServiceImpl struct {
	repo            repository.SetMealRepo
	setMealDishRepo repository.SetMealDishRepo
}

func (s *SetMealServiceImpl) OnOrClose(ctx context.Context, id uint64, status int) error {
	err := s.repo.SetStatus(ctx, id, status)
	return err
}

func (s *SetMealServiceImpl) PageQuery(ctx context.Context, dto request.SetMealPageQueryDTO) (*common.PageResult, error) {
	result, err := s.repo.PageQuery(ctx, dto)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SaveWithDish 保存套餐和菜品信息
func (s *SetMealServiceImpl) SaveWithDish(ctx context.Context, dto request.SetMealDTO) error {
	// 转换dto为model开启事务进行保存
	price, _ := strconv.ParseFloat(dto.Price, 10)
	setmeal := model.SetMeal{
		Id:          dto.Id,
		CategoryId:  dto.CategoryId,
		Name:        dto.Name,
		Price:       price,
		Status:      enum.ENABLE,
		Description: dto.Description,
		Image:       dto.Image,
	}
	// 开启事务进行存储
	transaction := s.repo.Transaction(ctx)
	defer func() {
		if err := recover(); err != nil {
			transaction.Rollback()
		}
	}()
	// 先插入套餐数据信息，并得到返回的主键id值
	err := s.repo.Insert(transaction, &setmeal)
	if err != nil {
		return err
	}
	// 再对中间表中套餐内的菜品信息附加主键id
	for _, setmealDish := range dto.SetMealDishs {
		setmealDish.SetmealId = setmeal.Id
	}
	// 向中间表插入数据
	err = s.setMealDishRepo.InsertBatch(transaction, dto.SetMealDishs)
	if err != nil {
		return err
	}
	return transaction.Commit().Error
}

func NewSetMealService(repo repository.SetMealRepo, setMealDishRepo repository.SetMealDishRepo) ISetMealService {
	return &SetMealServiceImpl{repo: repo, setMealDishRepo: setMealDishRepo}
}
