package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"take-out/common"
	"take-out/common/e"
	"take-out/common/enum"
	"take-out/global"
	"take-out/internal/api/request"
	"take-out/internal/service"
)

type EmployeeController struct {
	service service.IEmployeeService
}

func NewEmployeeController(employeeService service.IEmployeeService) *EmployeeController {
	return &EmployeeController{service: employeeService}
}

// Login 员工登录
func (ec *EmployeeController) Login(ctx *gin.Context) {
	code := e.SUCCESS
	employeeLogin := request.EmployeeLogin{}
	err := ctx.Bind(&employeeLogin)
	if err != nil {
		code = e.ERROR
		global.Log.Info("EmployeeController login 解析失败")
		return
	}
	resp, err := ec.service.Login(ctx, employeeLogin)
	if err != nil {
		code = e.ERROR
		global.Log.Info("EmployeeController login Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: resp,
	})

}

// Logout 员工退出
func (ec *EmployeeController) Logout(ctx *gin.Context) {
	code := e.SUCCESS
	var err error
	err = ec.service.Logout(ctx)
	if err != nil {
		code = e.ERROR
		global.Log.Info("EmployeeController login Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

// EditPassword 修改密码
func (ec *EmployeeController) EditPassword(ctx *gin.Context) {
	code := e.SUCCESS
	var reqs request.EmployeeEditPassword
	var err error
	err = ctx.Bind(&reqs)
	if err != nil {
		global.Log.Info("EditPassword Error:", err.Error())
		return
	}
	// 从上下文获取员工id
	if id, ok := ctx.Get(enum.CurrentId); ok {
		reqs.EmpId = id.(uint64)
	}
	err = ec.service.EditPassword(ctx, reqs)
	if err != nil {
		code = e.ERROR
		global.Log.Info("EditPassword  Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

// AddEmployee 新增员工
func (ec *EmployeeController) AddEmployee(ctx *gin.Context) {
	code := e.SUCCESS
	var request request.EmployeeDTO
	var err error
	err = ctx.Bind(&request)
	if err != nil {
		global.Log.Info("AddEmployee Error:", err.Error())
		return
	}
	err = ec.service.CreateEmployee(ctx, request)
	if err != nil {
		code = e.ERROR
		global.Log.Info("AddEmployee  Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
	}
	// 正确输出
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}
