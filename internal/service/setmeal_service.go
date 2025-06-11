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

type ISetMealService interface {
	SaveWithDish(ctx context.Context, dto request.SetMealDTO) error
	PageQuery(ctx context.Context, dto request.SetMealPageQueryDTO) (*common.PageResult, error)
	OnOrClose(ctx context.Context, id uint64, status int) error
	GetByIdWithDish(ctx context.Context, dishId uint64) (response.SetMealWithDishByIdVo, error)
	Delete(ctx context.Context, ids string) error
}

type SetMealServiceImpl struct {
	repo            *dao.SetMealDao
	setMealDishRepo *dao.SetMealDishDao
}

func (s *SetMealServiceImpl) Delete(ctx context.Context, ids string) error {
	idStrList := strings.Split(ids, ",")
	setMealIdList := make([]uint64, len(idStrList))
	for i, id := range idStrList {
		setMealIdList[i] = cast.ToUint64(id)
	}
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		//1.删除套餐和菜品中间表数据
		err := s.setMealDishRepo.WithTx(tx).DeleteBySetMealIds(ctx, setMealIdList...)
		if err != nil {
			global.Log.ErrContext(ctx, "SetMealServiceImpl.Delete failed err=%s", err.Error())
			return err
		}
		//2.删除对应套餐
		err = s.repo.WithTx(tx).DeleteByIds(ctx, setMealIdList...)
		if err != nil {
			global.Log.ErrContext(ctx, "SetMealServiceImpl.Delete failed err=%s", err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.ErrContext(ctx, "SetMealServiceImpl.Delete transaction failed err=%s", err.Error())
		return retcode.NewError(e.MysqlTransActionERR, e.GetMsg(e.MysqlTransActionERR))
	}
	return nil
}

func (s *SetMealServiceImpl) GetByIdWithDish(ctx context.Context, setMealId uint64) (response.SetMealWithDishByIdVo, error) {
	var setMealVo response.SetMealWithDishByIdVo
	// 获取事务
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 单独查询套餐
		setMeal, err := s.repo.WithTx(tx).GetByIdWithDish(ctx, setMealId)
		if err != nil {
			global.Log.ErrContext(ctx, "SetMealServiceImpl.GetByIdWithDish failed err=%s", err.Error())
			return err
		}
		// 查询中间表记录的套餐菜品信息
		dishList, err := s.setMealDishRepo.WithTx(tx).GetBySetMealId(ctx, setMealId)
		if err != nil {
			global.Log.ErrContext(ctx, "SetMealServiceImpl.GetByIdWithDish failed err=%s", err.Error())
			return err
		}
		// 数据组装
		setMealVo = response.SetMealWithDishByIdVo{
			Id:            setMeal.Id,
			CategoryId:    setMeal.CategoryId,
			CategoryName:  setMeal.Name,
			Description:   setMeal.Description,
			Image:         setMeal.Image,
			Name:          setMeal.Name,
			Price:         setMeal.Price,
			Status:        setMeal.Status,
			SetmealDishes: dishList,
			UpdateTime:    setMeal.UpdateTime,
		}
		return nil
	})
	if err != nil {
		global.Log.ErrContext(ctx, "SetMealServiceImpl.GetByIdWithDish transaction failed err=%s", err.Error())
		return setMealVo, retcode.NewError(e.MysqlTransActionERR, e.GetMsg(e.MysqlTransActionERR))
	}
	return setMealVo, err
}

func (s *SetMealServiceImpl) OnOrClose(ctx context.Context, id uint64, status int) error {
	err := s.repo.SetStatus(ctx, id, status)
	if err != nil {
		global.Log.ErrContext(ctx, "SetMealServiceImpl.OnOrClose failed err=%s", err.Error())
		return err
	}
	return nil
}

func (s *SetMealServiceImpl) PageQuery(ctx context.Context, dto request.SetMealPageQueryDTO) (*common.PageResult, error) {
	result, err := s.repo.PageQuery(ctx, dto)
	if err != nil {
		global.Log.ErrContext(ctx, "SetMealServiceImpl.PageQuery failed err=%s", err.Error())
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
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 先插入套餐数据信息，并得到返回的主键id值
		err := s.repo.WithTx(tx).Insert(ctx, &setmeal)
		if err != nil {
			global.Log.ErrContext(ctx, "SetMealServiceImpl.SaveWithDish failed err=%s", err.Error())
			return err
		}
		// 再对中间表中套餐内的菜品信息附加主键id
		for _, setmealDish := range dto.SetMealDishs {
			setmealDish.SetmealId = setmeal.Id
		}
		// 向中间表插入数据
		err = s.setMealDishRepo.WithTx(tx).InsertBatch(ctx, dto.SetMealDishs)
		if err != nil {
			global.Log.ErrContext(ctx, "SetMealServiceImpl.SaveWithDish failed err=%s", err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.ErrContext(ctx, "SetMealServiceImpl.SaveWithDish transaction failed err=%s", err.Error())
		return retcode.NewError(e.MysqlTransActionERR, e.GetMsg(e.MysqlTransActionERR))
	}
	return nil
}

func NewSetMealService(repo *dao.SetMealDao, setMealDishRepo *dao.SetMealDishDao) ISetMealService {
	return &SetMealServiceImpl{repo: repo, setMealDishRepo: setMealDishRepo}
}
