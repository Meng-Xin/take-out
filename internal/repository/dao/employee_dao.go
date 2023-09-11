package dao

import (
	"context"
	"gorm.io/gorm"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type EmployeeDao struct {
	db *gorm.DB
}

func NewEmployeeDao(db *gorm.DB) repository.EmployeeRepo {
	return &EmployeeDao{db: db}
}

func (e *EmployeeDao) GetByUserName(ctx context.Context, userName string) (*model.Employee, error) {
	var employee model.Employee
	err := e.db.WithContext(ctx).Where("username=?", userName).First(&employee).Error
	return &employee, err
}
