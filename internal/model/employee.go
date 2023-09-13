package model

import (
	"gorm.io/gorm"
	"take-out/common/enum"
	"time"
)

type Employee struct {
	Id         uint64
	Username   string
	Name       string
	Password   string
	Phone      string
	Sex        string
	IdNumber   string
	Status     int
	CreateTime time.Time
	UpdateTime time.Time
	CreateUser uint64
	UpdateUser uint64
}

func (e *Employee) BeforeCreate(tx *gorm.DB) error {
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

func (e *Employee) BeforeUpdate(tx *gorm.DB) error {
	// 在更新记录千自动填充更新时间
	e.UpdateTime = time.Now()
	// 从上下文获取用户信息
	value := tx.Statement.Context.Value(enum.CurrentId)
	if uid, ok := value.(uint64); ok {
		e.UpdateUser = uid
	}
	return nil
}
