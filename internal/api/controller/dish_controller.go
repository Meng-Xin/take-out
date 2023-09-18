package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

// PageQuery 菜品分页查询
func (dc *DishController) PageQuery(ctx *gin.Context) {
	code := e.SUCCESS
	var dishPageQueryDTO request.DishPageQueryDTO
	dishPageQueryDTO.CategoryId, _ = strconv.ParseUint(ctx.Query("categoryId"), 10, 64)
	dishPageQueryDTO.Page, _ = strconv.Atoi(ctx.Query("page"))
	dishPageQueryDTO.PageSize, _ = strconv.Atoi(ctx.Query("pageSize"))
	dishPageQueryDTO.Status, _ = strconv.Atoi(ctx.Query("status"))
	dishPageQueryDTO.Name = ctx.Query("name")
	//分页查询
	pageResult, err := dc.service.PageQuery(ctx, &dishPageQueryDTO)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("DishController PageQuery failed", "Error", err.Error())
		ctx.JSON(http.StatusOK, common.Result{Code: code, Msg: e.GetMsg(code)})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: pageResult,
	})
}
