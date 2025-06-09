package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"take-out/common/retcode"
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

func (sc *SetMealController) Delete(ctx *gin.Context) {
	ids := ctx.Query("ids")
	err := sc.service.Delete(ctx, ids)
	if err != nil {
		global.Log.Error("Delete 删除套餐失败！", "err", err)
		retcode.Fatal(ctx, err, "")
	}
	retcode.OK(ctx, "")
}

// SaveWithDish 保存套餐和菜品信息
func (sc *SetMealController) SaveWithDish(ctx *gin.Context) {
	var setmealDTO request.SetMealDTO
	err := ctx.Bind(&setmealDTO)
	if err != nil {
		global.Log.Debug("SaveWithDish保存套餐和菜品信息 结构体解析失败！", "Err:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	err = sc.service.SaveWithDish(ctx, setmealDTO)
	if err != nil {
		global.Log.Warn("SaveWithDish保存套餐和菜品信息！", "Err:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")
}

// PageQuery 套餐分页查询
func (sc *SetMealController) PageQuery(ctx *gin.Context) {
	// 数据组装
	var pageQueryDTO request.SetMealPageQueryDTO
	err := ctx.Bind(&pageQueryDTO)
	if err != nil {
		global.Log.Error("PageQuery invalid params err:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	// 分页查询
	result, err := sc.service.PageQuery(ctx, pageQueryDTO)
	if err != nil {
		global.Log.Warn("PageQuery 套餐分页查询失败！", "Err:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, result)
}

// OnOrClose 套餐启用禁用
func (sc *SetMealController) OnOrClose(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	status, _ := strconv.Atoi(ctx.Param("status"))
	// 设置套餐状态
	err := sc.service.OnOrClose(ctx, id, status)
	if err != nil {
		global.Log.Warn("OnOrClose 套餐启用禁用失败！", "Err:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")
}

// GetByIdWithDish 根据套餐id获取套餐和关联菜品信息
func (sc *SetMealController) GetByIdWithDish(ctx *gin.Context) {
	setMealId, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	// 获取套餐详情
	resp, err := sc.service.GetByIdWithDish(ctx, setMealId)
	if err != nil {
		global.Log.Warn("GetByIdWithDish 根据套餐id获取套餐和关联菜品信息", "Err:", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, resp)
}
