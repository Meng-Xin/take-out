package repository

import (
	"context"
	"take-out/internal/model"
)

type EmployeeRepo interface {
	// GetByUserName 根据username获取用户信息
	GetByUserName(ctx context.Context, userName string) (*model.Employee, error)
	// GetById 根据id获取用户信息
	GetById(ctx context.Context, id uint64) (*model.Employee, error)
	// UpData 动态修改
	UpData(ctx context.Context, employee model.Employee) error
	// Insert 插入员工数据
	Insert(ctx context.Context, entity model.Employee) error
}
