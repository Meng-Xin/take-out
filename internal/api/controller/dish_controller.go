package controller

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"strconv"
	"take-out/common/retcode"
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
	var dishDTO request.DishDTO
	err := ctx.Bind(&dishDTO)
	if err != nil {
		global.Log.Debug("param DishDTO failed", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	err = dc.service.AddDishWithFlavors(ctx, dishDTO)
	if err != nil {
		global.Log.Warn("AddDish failed", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")
}

// PageQuery 菜品分页查询
func (dc *DishController) PageQuery(ctx *gin.Context) {
	var dishPageQueryDTO request.DishPageQueryDTO
	err := ctx.Bind(&dishPageQueryDTO)
	if err != nil {
		global.Log.Error("DishController invalid params ", "Error", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	//分页查询
	pageResult, err := dc.service.PageQuery(ctx, &dishPageQueryDTO)
	if err != nil {
		global.Log.Error("DishController PageQuery failed", "Error", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, pageResult)
}

// GetById 根据id查询菜品信息
func (dc *DishController) GetById(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	// 根据id查询并获取口味数据
	dishVO, err := dc.service.GetByIdWithFlavors(ctx, id)
	if err != nil {
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, dishVO)
}

// List 根据分类id查询菜品信息
func (dc *DishController) List(ctx *gin.Context) {
	categoryId, _ := strconv.ParseUint(ctx.Query("categoryId"), 10, 64)
	// 根据id查询并获取口味数据
	dishVO, err := dc.service.List(ctx, categoryId)
	if err != nil {
		global.Log.Error("根据分类id查询菜品信息失败！", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, dishVO)
}

// OnOrClose 菜品启售或禁售
func (dc *DishController) OnOrClose(ctx *gin.Context) {

	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	status, _ := strconv.Atoi(ctx.Param("status"))
	// 根据id修改对应菜品的状态
	err := dc.service.OnOrClose(ctx, id, status)
	if err != nil {
		slog.Warn("菜品启售或禁售失败！", "Err:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")
}

// Update 修改菜品信息
func (dc *DishController) Update(ctx *gin.Context) {
	dishUpdateDTO := request.DishUpdateDTO{}
	err := ctx.Bind(&dishUpdateDTO)
	if err != nil {
		slog.Warn("Failed Json Bind", "Err:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	// 更新菜品以及菜品口味的关联数据
	err = dc.service.Update(ctx, dishUpdateDTO)
	if err != nil {
		slog.Warn("更新菜品信息失败！", "Err:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")
}

// Delete 删除菜品信息
func (dc *DishController) Delete(ctx *gin.Context) {
	ids := ctx.Query("ids")
	err := dc.service.Delete(ctx, ids)
	if err != nil {
		slog.Warn("删除菜品信息失败！", "Err:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")
}
