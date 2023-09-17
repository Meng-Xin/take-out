package admin

import (
	"github.com/gin-gonic/gin"
	"take-out/global"
	"take-out/internal/api/controller"
	"take-out/internal/repository/dao"
	"take-out/internal/service"
	"take-out/middle"
)

type CategoryRouter struct{}

func (cr *CategoryRouter) InitApiRouter(parent *gin.RouterGroup) {
	publicRouter := parent.Group("category")
	privateRouter := parent.Group("category")
	// 私有路由使用jwt验证
	privateRouter.Use(middle.VerifyJWTAdmin())
	// 依赖注入
	categoryCtrl := controller.NewCategoryController(
		service.NewCategoryService(dao.NewCategoryDao(global.DB)),
	)
	{
		publicRouter.POST("", categoryCtrl.AddCategory)
	}
}
