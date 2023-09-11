package service

import (
	"context"
	"take-out/common/e"
	"take-out/common/enum"
	"take-out/common/utils"
	"take-out/global"
	"take-out/internal/api/request"
	"take-out/internal/api/response"
	"take-out/internal/repository"
)

type IEmployeeService interface {
	Login(context.Context, request.EmployeeLogin) (*response.EmployeeLogin, error)
}

type EmployeeImpl struct {
	repo repository.EmployeeRepo
}

func NewEmployeeService(repo repository.EmployeeRepo) IEmployeeService {
	return &EmployeeImpl{repo: repo}
}

func (ei *EmployeeImpl) Login(ctx context.Context, employeeLogin request.EmployeeLogin) (*response.EmployeeLogin, error) {
	// 1.查询用户是否存在
	employee, err := ei.repo.GetByUserName(ctx, employeeLogin.UserName)
	if err != nil || employee == nil {
		return nil, e.Error_ACCOUNT_NOT_FOUND
	}
	// 2.校验密码
	password := utils.MD5V(employeeLogin.Password, "", 0)
	if password != employee.Password {
		return nil, e.Error_PASSWORD_ERROR
	}
	// 3.校验状态
	if employee.Status == enum.DISABLE {
		return nil, e.Error_ACCOUNT_LOCKED
	}
	// 生成Token
	token, _, err := utils.GenToken(employee.Id, employee.Name, []byte(global.Config.Jwt.Admin.Secret))
	if err != nil {
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
