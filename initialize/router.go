package initialize

import (
	"github.com/Meng-Xin/logger"
	"github.com/gin-gonic/gin"
	"take-out/global"
	"take-out/internal/router"
)

func routerInit() *gin.Engine {
	r := gin.Default()
	allRouter := router.AllRouter

	// 链路追踪日志中间件
	r.Use(logger.GinMiddleware(global.Log, "takeout"))

	// admin
	admin := r.Group("/admin")
	{
		allRouter.EmployeeRouter.InitApiRouter(admin)
		allRouter.CategoryRouter.InitApiRouter(admin)
		allRouter.DishRouter.InitApiRouter(admin)
		allRouter.CommonRouter.InitApiRouter(admin)
		allRouter.SetMealRouter.InitApiRouter(admin)
	}
	return r
}
