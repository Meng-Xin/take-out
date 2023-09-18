package model

import (
	"gorm.io/gorm"
	"take-out/common/enum"
	"time"
)

type Dish struct {
	Id          uint64    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name        string    `json:"name"`
	CategoryId  uint64    `json:"categoryId"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	CreateUser  uint64    `json:"createUser"`
	UpdateUser  uint64    `json:"updateUser"`
	// 一对多
	Flavors []DishFlavor `json:"flavors"`
}

func (e *Dish) BeforeCreate(tx *gorm.DB) error {
	// 自动填充 创建时间、创建人、更新时间、更新用户
	e.CreateTime = time.Now()
	e.UpdateTime = time.Now()
	// 从上下文获取用户信息
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		e.CreateUser = uid
		e.UpdateUser = uid
	}
	return nil
}

func (e *Dish) BeforeUpdate(tx *gorm.DB) error {
	// 在更新记录千自动填充更新时间
	e.UpdateTime = time.Now()
	// 从上下文获取用户信息
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		e.UpdateUser = uid
	}
	return nil
}

func (e *Dish) TableName() string {
	return "dish"
}
