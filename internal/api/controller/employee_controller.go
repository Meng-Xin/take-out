package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"take-out/common/enum"
	"take-out/common/retcode"
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
	employeeLogin := request.EmployeeLogin{}
	err := ctx.Bind(&employeeLogin)
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeController login 解析失败")
		retcode.Fatal(ctx, err, "")
		return
	}
	resp, err := ec.service.Login(ctx, employeeLogin)
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeController login Error: err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, resp)
}

// Logout 员工退出
func (ec *EmployeeController) Logout(ctx *gin.Context) {
	var err error
	err = ec.service.Logout(ctx)
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeController login Error: err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")
}

// EditPassword 修改密码
func (ec *EmployeeController) EditPassword(ctx *gin.Context) {
	var reqs request.EmployeeEditPassword
	var err error
	err = ctx.Bind(&reqs)
	if err != nil {
		global.Log.ErrContext(ctx, "EditPassword Error: err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	// 从上下文获取员工id
	if id, ok := ctx.Get(enum.CurrentId); ok {
		reqs.EmpId = id.(uint64)
	}
	err = ec.service.EditPassword(ctx, reqs)
	if err != nil {
		global.Log.ErrContext(ctx, "EditPassword  Error: err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
	}
	retcode.OK(ctx, "")
}

// AddEmployee 新增员工
func (ec *EmployeeController) AddEmployee(ctx *gin.Context) {
	var request request.EmployeeDTO
	var err error
	err = ctx.Bind(&request)
	if err != nil {
		global.Log.ErrContext(ctx, "AddEmployee Error: err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	err = ec.service.CreateEmployee(ctx, request)
	if err != nil {
		global.Log.ErrContext(ctx, "AddEmployee  Error: err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	// 正确输出
	retcode.OK(ctx, "")
}

// PageQuery 员工分页查询
func (ec *EmployeeController) PageQuery(ctx *gin.Context) {
	var employeePageQueryDTO request.EmployeePageQueryDTO
	err := ctx.Bind(&employeePageQueryDTO)
	if err != nil {
		global.Log.ErrContext(ctx, "AddEmployee  invalid params err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	// 进行分页查询
	pageResult, err := ec.service.PageQuery(ctx, employeePageQueryDTO)
	if err != nil {
		global.Log.ErrContext(ctx, "AddEmployee  Error: err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, pageResult)
}

// OnOrOff 启用Or禁用员工状态
func (ec *EmployeeController) OnOrOff(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	status, _ := strconv.Atoi(ctx.Param("status"))
	var err error
	err = ec.service.SetStatus(ctx, id, status)
	if err != nil {
		global.Log.ErrContext(ctx, "OnOrOff Status  Error: err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	// 更新员工状态
	global.Log.Info("启用Or禁用员工状态：", "id", id, "status", status)
	retcode.OK(ctx, "")
}

// UpdateEmployee 编辑员工信息
func (ec *EmployeeController) UpdateEmployee(ctx *gin.Context) {
	var employeeDTO request.EmployeeDTO
	err := ctx.Bind(&employeeDTO)
	if err != nil {
		global.Log.ErrContext(ctx, "UpdateEmployee Error: err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	// 修改员工信息
	err = ec.service.UpdateEmployee(ctx.Request.Context(), employeeDTO)
	if err != nil {
		global.Log.ErrContext(ctx, "UpdateEmployee Error: err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")
}

// GetById 获取员工信息根据id
func (ec *EmployeeController) GetById(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	employee, err := ec.service.GetById(ctx.Request.Context(), id)
	if err != nil {
		global.Log.ErrContext(ctx, "EmployeeCtrl GetById Error err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, employee)
}
