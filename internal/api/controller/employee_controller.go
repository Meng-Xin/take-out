package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"take-out/common"
	"take-out/common/e"
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

func (ec *EmployeeController) Login(ctx *gin.Context) {
	code := e.SUCCESS
	employeeLogin := request.EmployeeLogin{}
	err := ctx.Bind(&employeeLogin)
	if err != nil {
		code = e.ERROR
		global.Log.Info("EmployeeController login 解析失败")
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  e.GetMsg(code),
		})
		return
	}
	resp, err := ec.service.Login(ctx.Request.Context(), employeeLogin)
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
