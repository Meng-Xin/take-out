package request

import (
	"take-out/internal/model"
)

type DishDTO struct {
	Id          uint64             `json:"id"`
	Name        string             `json:"name"`
	CategoryId  uint64             `json:"categoryId"`
	Price       string             `json:"price"`
	Image       string             `json:"image"`
	Description string             `json:"description"`
	Status      int                `json:"status"`
	Flavors     []model.DishFlavor `json:"flavors"`
}

type DishPageQueryDTO struct {
	Page       int    `json:"page"`       // 分页查询的页数
	PageSize   int    `json:"pageSize"`   // 分页查询的页容量
	Name       string `json:"name"`       // 分页查询的name
	CategoryId uint64 `json:"categoryId"` // 分类ID:
	Status     int    `json:"status"`     // 菜品状态
}

type DishUpdateDTO struct {
	Id          uint64             `json:"id" `
	Name        string             `json:"name"`
	CategoryId  uint64             `json:"categoryId"`
	Price       string             `json:"price"`
	Image       string             `json:"image"`
	Description string             `json:"description"`
	Status      int                `json:"status"`
	Flavors     []model.DishFlavor `json:"flavors"`
}
