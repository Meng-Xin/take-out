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

type SetMealController struct {
	service service.ISetMealService
}

func NewSetMealController(service service.ISetMealService) *SetMealController {
	return &SetMealController{service: service}
}

// SaveWithDish 保存套餐和菜品信息
func (sc *SetMealController) SaveWithDish(ctx *gin.Context) {
	code := e.SUCCESS
	var setmealDTO request.SetMealDTO
	err := ctx.Bind(&setmealDTO)
	if err != nil {
		global.Log.Debug("SaveWithDish保存套餐和菜品信息 结构体解析失败！", "Err:", err.Error())
		e.Send(ctx, e.ERROR)
		return
	}
	err = sc.service.SaveWithDish(ctx, setmealDTO)
	if err != nil {
		global.Log.Warn("SaveWithDish保存套餐和菜品信息！", "Err:", err.Error())
		e.Send(ctx, e.ERROR)
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  e.GetMsg(code),
	})
}

// PageQuery 套餐分页查询
func (sc *SetMealController) PageQuery(ctx *gin.Context) {
	code := e.SUCCESS
	// 数据组装
	var pageQueryDTO request.SetMealPageQueryDTO
	pageQueryDTO.CategoryId, _ = strconv.ParseUint(ctx.Query("categoryId"), 10, 64)
	pageQueryDTO.Name = ctx.Query("name")
	pageQueryDTO.Status, _ = strconv.Atoi(ctx.Query("status"))
	pageQueryDTO.Page, _ = strconv.Atoi(ctx.Query("page"))
	pageQueryDTO.PageSize, _ = strconv.Atoi(ctx.Query("pageSize"))
	// 分页查询
	result, err := sc.service.PageQuery(ctx, pageQueryDTO)
	if err != nil {
		global.Log.Warn("PageQuery 套餐分页查询失败！", "Err:", err.Error())
		e.Send(ctx, e.ERROR)
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: result,
		Msg:  e.GetMsg(code),
	})
}

// OnOrClose 套餐启用禁用
func (sc *SetMealController) OnOrClose(ctx *gin.Context) {
	code := e.SUCCESS
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	status, _ := strconv.Atoi(ctx.Param("status"))
	// 设置套餐状态
	err := sc.service.OnOrClose(ctx, id, status)
	if err != nil {
		global.Log.Warn("OnOrClose 套餐启用禁用失败！", "Err:", err.Error())
		e.Send(ctx, e.ERROR)
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  e.GetMsg(code),
	})
}

// GetByIdWithDish 根据套餐id获取套餐和关联菜品信息
func (sc *SetMealController) GetByIdWithDish(ctx *gin.Context) {
	code := e.SUCCESS
	setMealId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	// 获取套餐详情
	resp, err := sc.service.GetByIdWithDish(ctx, setMealId)
	if err != nil {
		global.Log.Warn("GetByIdWithDish 根据套餐id获取套餐和关联菜品信息", "Err:", err.Error())
		e.Send(ctx, e.ERROR)
		return
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: resp,
		Msg:  e.GetMsg(code),
	})
}
