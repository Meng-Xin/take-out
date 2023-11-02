package response

import (
	"take-out/internal/model"
	"time"
)

type SetMealPageQueryVo struct {
	Id           uint64    `json:"id" `          // 主键id
	CategoryId   uint64    `json:"categoryId"`   // 分类id
	Name         string    `json:"name"`         // 套餐名称
	Price        float64   `json:"price"`        // 套餐单价
	Status       int       `json:"status"`       // 套餐状态
	Description  string    `json:"description"`  // 套餐描述
	Image        string    `json:"image"`        // 套餐图片
	CreateTime   time.Time `json:"createTime"`   // 创建时间
	UpdateTime   time.Time `json:"updateTime"`   // 更新时间
	CreateUser   uint64    `json:"createUser"`   // 创建用户
	UpdateUser   uint64    `json:"updateUser"`   // 更新用户
	CategoryName string    `json:"categoryName"` // 分类名称
}

type SetMealWithDishByIdVo struct {
	Id            uint64              `json:"id"`
	CategoryId    uint64              `json:"categoryId"`
	CategoryName  string              `json:"categoryName"`
	Description   string              `json:"description"`
	Image         string              `json:"image"`
	Name          string              `json:"name"`
	Price         float64             `json:"price"`
	SetmealDishes []model.SetMealDish `json:"setmealDishes"`
	Status        int                 `json:"status"`
	UpdateTime    time.Time           `json:"updateTime"`
}
