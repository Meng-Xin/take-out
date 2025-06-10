package dao

import (
	"context"
	"gorm.io/gorm"
	"take-out/common"
	"take-out/common/e"
	"take-out/common/retcode"
	"take-out/global"
	"take-out/internal/api/request"
	"take-out/internal/model"
)

type EmployeeDao struct {
	db *gorm.DB
}

// UpdateStatus 动态更新包括零值
func (d *EmployeeDao) UpdateStatus(ctx context.Context, employee model.Employee) error {
	err := d.db.WithContext(ctx).Model(&model.Employee{}).Where("id = ?",
		employee.Id).Update("status", employee.Status).Error
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeDao.UpdateStatus failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "update employee failed")
	}
	return nil
}

func (d *EmployeeDao) PageQuery(ctx context.Context, dto request.EmployeePageQueryDTO) (*common.PageResult, error) {
	// 分页查询 select count(*) from employee where name = ? limit x,y
	var result common.PageResult
	var employeeList []model.Employee
	var err error
	// 动态拼接
	query := d.db.WithContext(ctx).Model(&model.Employee{})
	if dto.Name != "" {
		query = query.Where("name LIKE ?", "%"+dto.Name+"%")
	}
	// 计算总数
	if err = query.Count(&result.Total).Error; err != nil {
		global.Log.ErrContext(ctx, "EmployeeDao.PageQuery Count failed, err: %v", err)
		return nil, retcode.NewError(e.MysqlERR, "Get employee List failed")
	}
	// 分页查询
	err = query.Scopes(result.Paginate(&dto.Page, &dto.PageSize)).Find(&employeeList).Error
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeDao.PageQuery List failed, err: %v", err)
		return nil, retcode.NewError(e.MysqlERR, "Get employee List failed")
	}
	result.Records = employeeList
	return &result, nil
}

func (d *EmployeeDao) Insert(ctx context.Context, entity model.Employee) error {
	err := d.db.WithContext(ctx).Create(&entity).Error
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeDao.Insert failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "Create employee failed")
	}
	return nil
}

func (d *EmployeeDao) GetById(ctx context.Context, id uint64) (*model.Employee, error) {
	var employee model.Employee
	err := d.db.WithContext(ctx).Where("id=?", id).First(&employee).Error
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeDao.GetById failed, err: %v", err)
		return nil, retcode.NewError(e.MysqlERR, "Get employee failed")
	}
	return &employee, nil
}

func (d *EmployeeDao) Update(ctx context.Context, employee model.Employee) error {
	err := d.db.WithContext(ctx).Model(&employee).Updates(employee).Error
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeDao.Update failed, err: %v", err)
		return retcode.NewError(e.MysqlERR, "Update employee failed")
	}
	return nil
}

func (d *EmployeeDao) GetByUserName(ctx context.Context, userName string) (*model.Employee, error) {
	var employee model.Employee
	err := d.db.WithContext(ctx).Where("username=?", userName).First(&employee).Error
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeDao.GetByUserName failed, err: %v", err)
		return nil, retcode.NewError(e.MysqlERR, "Get employee failed")
	}
	return &employee, nil
}

func NewEmployeeDao(db *gorm.DB) *EmployeeDao {
	return &EmployeeDao{db: db}
}
