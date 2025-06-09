package admin

import (
	"github.com/gin-gonic/gin"
	"take-out/global"
	"take-out/internal/api/controller"
	"take-out/internal/repository/dao"
	"take-out/internal/service"
	"take-out/middle"
)

type SetMealRouter struct{}

func (sr *SetMealRouter) InitApiRouter(parent *gin.RouterGroup) {
	//publicRouter := parent.Group("category")
	privateRouter := parent.Group("setmeal")
	// 私有路由使用jwt验证
	privateRouter.Use(middle.VerifyJWTAdmin())
	// 依赖注入
	setmealCtrl := controller.NewSetMealController(
		service.NewSetMealService(dao.NewSetMealDao(global.DB), dao.NewSetMealDishDao()),
	)
	{
		privateRouter.POST("", setmealCtrl.SaveWithDish)
		privateRouter.GET("/page", setmealCtrl.PageQuery)
		privateRouter.GET("/:id", setmealCtrl.GetByIdWithDish)
		privateRouter.POST("/status/:status", setmealCtrl.OnOrClose)
		privateRouter.DELETE("", setmealCtrl.Delete)
	}
}
