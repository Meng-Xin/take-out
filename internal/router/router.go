package router

import "take-out/internal/router/admin"

type RouterGroup struct {
	admin.EmployeeRouter
	admin.CategoryRouter
}

var AllRouter = new(RouterGroup)
