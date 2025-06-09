package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"strconv"
	"take-out/common/retcode"
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
	var categoryDto request.CategoryDTO
	err := ctx.Bind(&categoryDto)
	if err != nil {
		global.Log.DebugContext(ctx, "param CategoryDTO json failed err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	if err = cc.service.AddCategory(ctx, categoryDto); err != nil {
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")
}

func (cc *CategoryController) PageQuery(ctx *gin.Context) {
	var categoryPageDTO request.CategoryPageQueryDTO
	err := ctx.Bind(&categoryPageDTO)
	if err != nil {
		global.Log.ErrContext(ctx, "Category invalid params failed err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	query, err := cc.service.PageQuery(ctx, categoryPageDTO)
	if err != nil {
		global.Log.ErrContext(ctx, "Category PageQuery failed err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, query)
}

func (cc *CategoryController) List(ctx *gin.Context) {
	cate, _ := strconv.Atoi(ctx.Query("type"))
	list, err := cc.service.List(ctx, cate)
	if err != nil {
		global.Log.ErrContext(ctx, "Category List failed err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, list)
}

func (cc *CategoryController) DeleteById(ctx *gin.Context) {
	id := cast.ToUint64(ctx.Query("id"))
	err := cc.service.DeleteById(ctx, id)
	if err != nil {
		global.Log.ErrContext(ctx, "Category DeleteById failed err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")
}

func (cc *CategoryController) EditCategory(ctx *gin.Context) {
	var categoryDTO request.CategoryDTO
	err := ctx.Bind(&categoryDTO)
	if err != nil {
		global.Log.ErrContext(ctx, "param CategoryDTO failed err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	err = cc.service.Update(ctx, categoryDTO)
	if err != nil {
		global.Log.ErrContext(ctx, "Category Edit failed err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")

}

func (cc *CategoryController) SetStatus(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Query("id"), 10, 64)
	status, _ := strconv.Atoi(ctx.Param("status"))
	err := cc.service.SetStatus(ctx, id, status)
	if err != nil {
		global.Log.ErrContext(ctx, "Category SetStatus failed err=%s", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, "")
}
