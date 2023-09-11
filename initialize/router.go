package initialize

import (
	"github.com/gin-gonic/gin"
	"take-out/internal/router"
)

func routerInit() *gin.Engine {
	r := gin.Default()
	allrouter := router.AllRouter
	// Swagger

	// admin
	admin := r.Group("/admin")
	{
		allrouter.EmployeeRouter.InitApiRouter(admin)
	}
	return r
}
