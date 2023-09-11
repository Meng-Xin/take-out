package repository

import (
	"context"
	"take-out/internal/model"
)

type EmployeeRepo interface {
	// GetByUserName 根据username获取用户信息
	GetByUserName(ctx context.Context, userName string) (*model.Employee, error)
}
