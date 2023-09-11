package admin

import (
	"github.com/gin-gonic/gin"
	"take-out/global"
	"take-out/internal/api/controller"
	"take-out/internal/repository/dao"
	"take-out/internal/service"
)

type EmployeeRouter struct{ service service.IEmployeeService }

func (er *EmployeeRouter) InitApiRouter(router *gin.RouterGroup) {
	employeeRouter := router.Group("employee")
	//employeeRouter.Use(middle.VerifyJWTAdmin())
	// 依赖注入
	er.service = service.NewEmployeeService(
		dao.NewEmployeeDao(global.DB),
	)
	employeeCtl := controller.NewEmployeeController(er.service)
	{
		employeeRouter.POST("/login", employeeCtl.Login)
	}
}
