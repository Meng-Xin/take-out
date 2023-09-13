package dao

import (
	"context"
	"gorm.io/gorm"
	"take-out/common"
	"take-out/internal/api/request"
	"take-out/internal/model"
	"take-out/internal/repository"
)

type EmployeeDao struct {
	db *gorm.DB
}

// UpDataZero 动态更新包括零值
func (e *EmployeeDao) UpdateStatus(ctx context.Context, employee model.Employee) error {
	err := e.db.WithContext(ctx).Model(&model.Employee{}).Where("id = ?", employee.Id).Update("status", employee.Status).Error
	return err
}

func (e *EmployeeDao) PageQuery(ctx context.Context, dto request.EmployeePageQueryDTO) (*common.PageResult, error) {
	// 分页查询 select count(*) from employee where name = ? limit x,y
	var result common.PageResult
	var employeeList []model.Employee
	var err error
	// 动态拼接
	query := e.db.WithContext(ctx).Model(&model.Employee{})
	if dto.Name != "" {
		query = query.Where("name LIKE ?", "%"+dto.Name+"%")
	}
	// 计算总数
	if err = query.Count(&result.Total).Error; err != nil {
		return nil, err
	}
	// 分页查询
	err = query.Scopes(result.Paginate(&dto.Page, &dto.PageSize)).Find(&employeeList).Error
	result.Records = employeeList
	return &result, err
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

func (e *EmployeeDao) Update(ctx context.Context, employee model.Employee) error {
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
