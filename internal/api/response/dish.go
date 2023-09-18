package response

import (
	"take-out/internal/model"
	"time"
)

type DishPageVo struct {
	Id           uint64    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name         string    `json:"name"`
	CategoryId   uint64    `json:"categoryId"`
	Price        float64   `json:"price"`
	Image        string    `json:"image"`
	Description  string    `json:"description"`
	Status       int       `json:"status"`
	UpdateTime   time.Time `json:"updateTime"`
	CategoryName string    `json:"categoryName"`
}

type DishVo struct {
	Id          uint64    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name        string    `json:"name"`
	CategoryId  uint64    `json:"categoryId"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	UpdateTime  time.Time `json:"updateTime"`
	//CategoryName string             `json:"categoryName"`
	Flavors []model.DishFlavor `json:"flavors"`
}
