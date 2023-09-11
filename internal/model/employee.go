package model

import (
	"gorm.io/gorm"
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
	// 在插入记录千自动填充创建时间
	e.CreateTime = time.Now()
	return nil
}

func (e *Employee) BeforeUpdate(tx *gorm.DB) error {
	// 在更新记录千自动填充更新时间
	e.UpdateTime = time.Now()
	return nil
}
