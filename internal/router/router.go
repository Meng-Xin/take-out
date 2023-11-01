package router

import "take-out/internal/router/admin"

type RouterGroup struct {
	admin.EmployeeRouter
	admin.CategoryRouter
	admin.DishRouter
	admin.CommonRouter
	admin.SetMealRouter
}

var AllRouter = new(RouterGroup)
