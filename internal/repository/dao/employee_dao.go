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

func (e *EmployeeDao) Insert(ctx context.Context, entity model.Employee) error {
	err := e.db.WithContext(ctx).Create(&entity).Error
	return err
}

func (e *EmployeeDao) GetById(ctx context.Context, id uint64) (*model.Employee, error) {
	var employee model.Employee
	err := e.db.WithContext(ctx).Where("id=?", id).First(&employee).Error
	return &employee, err
}

func (e *EmployeeDao) UpData(ctx context.Context, employee model.Employee) error {
	err := e.db.WithContext(ctx).Model(&employee).Updates(employee).Error
	return err
}

func (e *EmployeeDao) GetByUserName(ctx context.Context, userName string) (*model.Employee, error) {
	var employee model.Employee
	err := e.db.WithContext(ctx).Where("username=?", userName).First(&employee).Error
	return &employee, err
}

func NewEmployeeDao(db *gorm.DB) repository.EmployeeRepo {
	return &EmployeeDao{db: db}
}
