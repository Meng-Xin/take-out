package e

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"take-out/common"
)

const (
	SUCCESS                         = 1    //ok
	ERROR                           = 2    //内部错误
	UNKNOW_IDENTITY                 = 403  //未知身份
	MysqlERR                        = 1001 //mysql出错
	MysqlTransActionERR             = 1002 //mysql事务执行出错
	RedisERR                        = 1003 //redis出错
	ErrorPasswordError              = 2001 //密码错误
	ErrorAccountNotFound            = 2002 //账号不存在
	ErrorAccountLOCKED              = 2003 //账号被锁定
	ErrorUnknowError                = 2004 //未知错误
	ErrorUserNotLogin               = 2005 //用户未登录
	ErrorCategoryBeRelatedBySetmeal = 2006 //当前分类关联了套餐,不能删除
	ErrorCategoryBeRelatedByDish    = 2007 //当前分类关联了菜品,不能删除
	ErrorShoppingCartIsNull         = 2008 //购物车数据为空，不能下单
	ErrorAddressBookIsNull          = 2009 //地址为空，不能下单
	ErrorLoginFailed                = 2010 //登录失败
	ErrorUploadFailed               = 2011 //文件上传失败
	ErrorSetMealEnableFailed        = 2012 //套餐内包含未启售菜品，无法启售
	ErrorPasswordEditFailed         = 2013 //密码修改失败
	ErrorDishOnSale                 = 2014 //起售中的菜品不能删除
	ErrorSetMEALOnSale              = 2015 //起售中的套餐不能删除
	ErrorDishBeRelatedBySetMeal     = 2016 //当前菜品关联了套餐,不能删除
	ErrorDishBeRelatedByOrder       = 2017 //当前菜品关联了订单,不能删除
	ErrorOrderStatusError           = 2018 //订单状态错误
	ErrorOrderNotFound              = 2019 //订单不存在

)

var ErrMsg = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "内部错误",
	UNKNOW_IDENTITY:                 "未知身份",
	MysqlTransActionERR:             "内部错误",
	ErrorPasswordError:              "密码错误",
	ErrorAccountNotFound:            "账号不存在",
	ErrorAccountLOCKED:              "账号被锁定",
	ErrorUnknowError:                "未知错误",
	ErrorUserNotLogin:               "用户未登录",
	ErrorCategoryBeRelatedBySetmeal: "当前分类关联了套餐,不能删除",
	ErrorCategoryBeRelatedByDish:    "当前分类关联了菜品,不能删除",
	ErrorShoppingCartIsNull:         "购物车数据为空，不能下单",
	ErrorAddressBookIsNull:          "地址为空，不能下单",
	ErrorLoginFailed:                "登录失败",
	ErrorUploadFailed:               "文件上传失败",
	ErrorSetMealEnableFailed:        "套餐内包含未启售菜品，无法启售",
	ErrorPasswordEditFailed:         "密码修改失败",
	ErrorDishOnSale:                 "起售中的菜品不能删除",
	ErrorSetMEALOnSale:              "起售中的套餐不能删除",
	ErrorDishBeRelatedBySetMeal:     "当前菜品关联了套餐,不能删除",
	ErrorDishBeRelatedByOrder:       "当前菜品关联了订单,不能删除",
	ErrorOrderStatusError:           "订单状态错误",
	ErrorOrderNotFound:              "订单不存在",
}

func GetMsg(code int) string {
	return ErrMsg[code]
}

func Send(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusOK, common.Result{Code: code, Msg: GetMsg(code)})
}
