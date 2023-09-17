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
	if err = cc.service.AddCategory(ctx, categoryDto); err != nil {
		code = e.ERROR
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (cc *CategoryController) PageQuery(ctx *gin.Context) {
	code := e.SUCCESS
	var categoryPageDTO request.CategoryPageQueryDTO
	categoryPageDTO.Name = ctx.Query("name")
	categoryPageDTO.Page, _ = strconv.Atoi(ctx.Query("page"))
	categoryPageDTO.PageSize, _ = strconv.Atoi(ctx.Query("pageSize"))
	categoryPageDTO.Cate, _ = strconv.Atoi(ctx.Query("type"))
	query, err := cc.service.PageQuery(ctx, categoryPageDTO)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("Category PageQuery failed ", err.Error())
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: query,
	})
}

func (cc *CategoryController) List(ctx *gin.Context) {
	code := e.SUCCESS
	cate, _ := strconv.Atoi(ctx.Query("type"))
	list, err := cc.service.List(ctx, cate)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("Category List failed", err.Error())
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: list,
	})
}

func (cc *CategoryController) DeleteById(ctx *gin.Context) {
	code := e.SUCCESS
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	err := cc.service.DeleteById(ctx, id)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("Category DeleteById failed", err.Error())
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (cc *CategoryController) EditCategory(ctx *gin.Context) {
	code := e.SUCCESS
	var categoryDTO request.CategoryDTO
	err := ctx.Bind(&categoryDTO)
	if err != nil {
		code = e.ERROR
		global.Log.Debug("param CategoryDTO failed", err.Error())
		return
	}
	err = cc.service.Update(ctx, categoryDTO)
	if err != nil {
		code = e.ERROR
		global.Log.Debug("Category Edit failed", err.Error())
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
	})
}

func (cc *CategoryController) SetStatus(ctx *gin.Context) {
	code := e.SUCCESS
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	status, _ := strconv.Atoi(ctx.Param("status"))
	err := cc.service.SetStatus(ctx, id, status)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("Category SetStatus failed", err.Error())
	}
	ctx.JSON(http.StatusOK, common.Result{Code: code, Msg: e.GetMsg(code)})
}
