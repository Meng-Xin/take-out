package controller

import (
	"github.com/gin-gonic/gin"
	"log/slog"
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

// GetById 根据id查询菜品信息
func (dc *DishController) GetById(ctx *gin.Context) {
	code := e.SUCCESS
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	// 根据id查询并获取口味数据
	dishVO, err := dc.service.GetByIdWithFlavors(ctx, id)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: dishVO,
	})
}

// List 根据分类id查询菜品信息
func (dc *DishController) List(ctx *gin.Context) {
	code := e.SUCCESS
	categoryId, _ := strconv.ParseUint(ctx.Query("categoryId"), 10, 64)
	// 根据id查询并获取口味数据
	dishVO, err := dc.service.List(ctx, categoryId)
	if err != nil {
		code = e.ERROR
		slog.Warn("根据分类id查询菜品信息失败！", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{Code: code, Msg: e.GetMsg(code)})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: dishVO,
	})
}

// OnOrClose 菜品启售或禁售
func (dc *DishController) OnOrClose(ctx *gin.Context) {
	code := e.SUCCESS
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	status, _ := strconv.Atoi(ctx.Param("status"))
	// 根据id修改对应菜品的状态
	err := dc.service.OnOrClose(ctx, id, status)
	if err != nil {
		code = e.ERROR
		slog.Warn("菜品启售或禁售失败！", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{Code: code, Msg: e.GetMsg(code)})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  e.GetMsg(code),
	})
}

// Update 修改菜品信息
func (dc *DishController) Update(ctx *gin.Context) {
	code := e.SUCCESS
	dishUpdateDTO := request.DishUpdateDTO{}
	err := ctx.Bind(&dishUpdateDTO)
	if err != nil {
		code = e.ERROR
		slog.Warn("Failed Json Bind", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{Code: code, Msg: e.GetMsg(code)})
		return
	}
	// 更新菜品以及菜品口味的关联数据
	err = dc.service.Update(ctx, dishUpdateDTO)
	if err != nil {
		code = e.ERROR
		slog.Warn("更新菜品信息失败！", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{Code: code, Msg: e.GetMsg(code)})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  e.GetMsg(code),
	})
}

// Delete 删除菜品信息
func (dc *DishController) Delete(ctx *gin.Context) {
	code := e.SUCCESS
	ids := ctx.Query("ids")
	err := dc.service.Delete(ctx, ids)
	if err != nil {
		code = e.ERROR
		slog.Warn("删除菜品信息失败！", "Err:", err.Error())
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  e.GetMsg(code),
		})
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  e.GetMsg(code),
	})
}
