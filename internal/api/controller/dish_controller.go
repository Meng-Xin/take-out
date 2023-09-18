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

type DishController struct {
	service service.IDishService
}

func NewDishController(service service.IDishService) *DishController {
	return &DishController{service: service}
}

// AddDish 新增菜品数据
func (dc *DishController) AddDish(ctx *gin.Context) {
	code := e.SUCCESS
	var dishDTO request.DishDTO
	err := ctx.Bind(&dishDTO)
	if err != nil {
		code = e.ERROR
		global.Log.Debug("param DishDTO failed", err.Error())
		return
	}
	err = dc.service.AddDishWithFlavors(ctx, dishDTO)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("AddDish failed", err.Error())
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  e.GetMsg(code),
	})
}
