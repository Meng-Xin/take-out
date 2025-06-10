package service

import (
	"context"
	"take-out/common"
	"take-out/common/e"
	"take-out/common/enum"
	"take-out/common/retcode"
	"take-out/common/utils"
	"take-out/global"
	"take-out/internal/api/request"
	"take-out/internal/api/response"
	"take-out/internal/model"
	"take-out/internal/repository/dao"
)

type IEmployeeService interface {
	Login(context.Context, request.EmployeeLogin) (*response.EmployeeLogin, error)
	Logout(ctx context.Context) error
	EditPassword(context.Context, request.EmployeeEditPassword) error
	CreateEmployee(ctx context.Context, employee request.EmployeeDTO) error
	PageQuery(ctx context.Context, dto request.EmployeePageQueryDTO) (*common.PageResult, error)
	SetStatus(ctx context.Context, id uint64, status int) error
	UpdateEmployee(ctx context.Context, dto request.EmployeeDTO) error
	GetById(ctx context.Context, id uint64) (*model.Employee, error)
}

type EmployeeImpl struct {
	repo *dao.EmployeeDao
}

func (ei *EmployeeImpl) GetById(ctx context.Context, id uint64) (*model.Employee, error) {
	employee, err := ei.repo.GetById(ctx, id)
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeImpl.GetById failed, err: %v", err)
		return nil, err
	}
	employee.Password = "***"
	return employee, err
}

func (ei *EmployeeImpl) UpdateEmployee(ctx context.Context, dto request.EmployeeDTO) error {
	// 构建model实体进行更新
	err := ei.repo.Update(ctx, model.Employee{
		Id:       dto.Id,
		Username: dto.UserName,
		Name:     dto.Name,
		Phone:    dto.Phone,
		Sex:      dto.Sex,
		IdNumber: dto.IdNumber,
	})
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeImpl.UpdateEmployee failed, err: %v", err)
		return err
	}
	return nil
}

func (ei *EmployeeImpl) SetStatus(ctx context.Context, id uint64, status int) error {
	// 设置用户状态,构造实体
	entity := model.Employee{Id: id, Status: status}
	err := ei.repo.UpdateStatus(ctx, entity)
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeImpl.SetStatus failed, err: %v", err)
		return err
	}
	return nil
}

func (ei *EmployeeImpl) PageQuery(ctx context.Context, dto request.EmployeePageQueryDTO) (*common.PageResult, error) {
	// 分页查询
	pageResult, err := ei.repo.PageQuery(ctx, dto)
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeImpl.PageQuery failed, err: %v", err)
		return nil, err
	}
	// 屏蔽敏感信息
	if employees, ok := pageResult.Records.([]model.Employee); ok {
		// 替换敏感信息
		for key, _ := range employees {
			employees[key].Password = "****"
			employees[key].IdNumber = "****"
			employees[key].Phone = "****"
		}
		// 重新赋值
		pageResult.Records = employees
	}

	return pageResult, nil
}

func (ei *EmployeeImpl) CreateEmployee(ctx context.Context, employeeDTO request.EmployeeDTO) error {
	var err error
	// 1.新增员工,构建员工基础信息
	entity := model.Employee{
		Id:       employeeDTO.Id,
		IdNumber: employeeDTO.IdNumber,
		Name:     employeeDTO.Name,
		Phone:    employeeDTO.Phone,
		Sex:      employeeDTO.Sex,
		Username: employeeDTO.UserName,
	}
	// 新增用户为启用状态
	entity.Status = enum.ENABLE
	// 新增用户初始密码为123456
	entity.Password = utils.MD5V("123456", "", 0)
	// 新增用户
	err = ei.repo.Insert(ctx, entity)
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeImpl.CreateEmployee failed, err: %v", err)
		return err
	}
	return nil
}

func (ei *EmployeeImpl) EditPassword(ctx context.Context, employeeEdit request.EmployeeEditPassword) error {
	// 1.获取员工信息
	employee, err := ei.repo.GetById(ctx, employeeEdit.EmpId)
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeImpl.EditPassword failed, err: %v", err)
		return err
	}
	// 校验用户老密码
	if employee == nil {
		return retcode.NewError(e.ErrorAccountNotFound, e.GetMsg(e.ErrorAccountNotFound))
	}
	oldHashPassword := utils.MD5V(employeeEdit.OldPassword, "", 0)
	if employee.Password != oldHashPassword {
		return retcode.NewError(e.ErrorPasswordError, e.GetMsg(e.ErrorPasswordError))
	}
	// 修改员工密码
	newHashPassword := utils.MD5V(employeeEdit.NewPassword, "", 0) // 使用新密码生成哈希值
	err = ei.repo.Update(ctx, model.Employee{
		Id:       employeeEdit.EmpId,
		Password: newHashPassword,
	})
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeImpl.EditPassword failed, err: %v", err)
		return err
	}
	return nil
}

func (ei *EmployeeImpl) Logout(ctx context.Context) error {
	// TODO 后续扩展为单点登录模式。 1.获取上下文中当前用户
	// 2.如果是单点登录的话执行推出操作
	return nil
}

func (ei *EmployeeImpl) Login(ctx context.Context, employeeLogin request.EmployeeLogin) (*response.EmployeeLogin, error) {
	// 1.查询用户是否存在
	employee, err := ei.repo.GetByUserName(ctx, employeeLogin.UserName)
	if err != nil || employee == nil {
		return nil, retcode.NewError(e.ErrorAccountNotFound, e.GetMsg(e.ErrorAccountNotFound))
	}
	// 2.校验密码
	password := utils.MD5V(employeeLogin.Password, "", 0)
	if password != employee.Password {
		return nil, retcode.NewError(e.ErrorPasswordError, e.GetMsg(e.ErrorPasswordError))
	}
	// 3.校验状态
	if employee.Status == enum.DISABLE {
		return nil, retcode.NewError(e.ErrorAccountLOCKED, e.GetMsg(e.ErrorAccountLOCKED))
	}
	// 生成Token
	jwtConfig := global.Config.Jwt.Admin
	token, err := utils.GenerateToken(employee.Id, jwtConfig.Name, jwtConfig.Secret)
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeImpl.Login failed, err: %v", err)
		return nil, err
	}
	// 4.构造返回数据
	resp := response.EmployeeLogin{
		Id:       employee.Id,
		Name:     employee.Name,
		Token:    token,
		UserName: employee.Username,
	}
	return &resp, nil
}

func NewEmployeeService(repo *dao.EmployeeDao) IEmployeeService {
	return &EmployeeImpl{repo: repo}
}
