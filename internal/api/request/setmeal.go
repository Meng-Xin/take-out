package request

import "take-out/internal/model"

type SetMealDTO struct {
	Id           uint64              `json:"id"`            // 主键id
	CategoryId   uint64              `json:"categoryId"`    // 分类id
	Name         string              `json:"name"`          // 套餐名称
	Price        string              `json:"price"`         // 套餐单价
	Status       int                 `json:"status"`        // 套餐状态
	Description  string              `json:"description"`   // 套餐描述
	Image        string              `json:"image"`         // 套餐图片
	SetMealDishs []model.SetMealDish `json:"setmealDishes"` // 套餐菜品关系
}

type SetMealPageQueryDTO struct {
	Page       int    `form:"page"`       // 分页查询的页数
	PageSize   int    `form:"pageSize"`   // 分页查询的页容量
	Name       string `form:"name"`       // 分页查询的name
	CategoryId uint64 `form:"categoryId"` // 分类ID:
	Status     int    `form:"status"`     // 套餐起售状态
}
