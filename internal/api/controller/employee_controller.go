package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		global.Log.Debug("EmployeeController login 解析失败")
		return
	}
	resp, err := ec.service.Login(ctx, employeeLogin)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("EmployeeController login Error:", err.Error())
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
		global.Log.Warn("EmployeeController login Error:", err.Error())
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
		global.Log.Debug("EditPassword Error:", err.Error())
		return
	}
	// 从上下文获取员工id
	if id, ok := ctx.Get(enum.CurrentId); ok {
		reqs.EmpId = id.(uint64)
	}
	err = ec.service.EditPassword(ctx, reqs)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("EditPassword  Error:", err.Error())
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
		global.Log.Debug("AddEmployee Error:", err.Error())
		return
	}
	err = ec.service.CreateEmployee(ctx, request)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("AddEmployee  Error:", err.Error())
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

// PageQuery 员工分页查询
func (ec *EmployeeController) PageQuery(ctx *gin.Context) {
	code := e.SUCCESS
	var employeePageQueryDTO request.EmployeePageQueryDTO
	employeePageQueryDTO.Name = ctx.Query("name")
	employeePageQueryDTO.Page, _ = strconv.Atoi(ctx.Query("page"))
	employeePageQueryDTO.PageSize, _ = strconv.Atoi(ctx.Query("pageSize"))
	// 进行分页查询
	pageResult, err := ec.service.PageQuery(ctx, employeePageQueryDTO)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("AddEmployee  Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: pageResult,
	})
}

// OnOrOff 启用Or禁用员工状态
func (ec *EmployeeController) OnOrOff(ctx *gin.Context) {
	code := e.SUCCESS
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	status, _ := strconv.Atoi(ctx.Param("status"))
	var err error
	err = ec.service.SetStatus(ctx, id, status)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("OnOrOff Status  Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
	}
	// 更新员工状态
	global.Log.Info("启用Or禁用员工状态：", "id", id, "status", status)
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

// UpdateEmployee 编辑员工信息
func (ec *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	code := e.SUCCESS
	var employeeDTO request.EmployeeDTO
	err := ctx.Bind(&employeeDTO)
	if err != nil {
		global.Log.Debug("UpdateEmployee Error:", err.Error())
		return
	}
	// 修改员工信息
	err = ec.service.UpdateEmployee(ctx.Request.Context(), employeeDTO)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("UpdateEmployee Error:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

// GetById 获取员工信息根据id
func (ec *EmployeeController) GetById(ctx *gin.Context) {
	code := e.SUCCESS
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	employee, err := ec.service.GetById(ctx.Request.Context(), id)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("EmployeeCtrl GetById Error", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: employee,
	})
}
