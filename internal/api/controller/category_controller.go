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

type CategoryController struct {
	service service.ICategoryService
}

func NewCategoryController(service service.ICategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (cc *CategoryController) AddCategory(ctx *gin.Context) {
	code := e.SUCCESS
	var categoryDto request.CategoryDTO
	err := ctx.Bind(&categoryDto)
	if err != nil {
		global.Log.Debug("param CategoryDTO json failed", err.Error())
		return
	}
	err = cc.service.AddCategory(ctx, categoryDto)
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}
