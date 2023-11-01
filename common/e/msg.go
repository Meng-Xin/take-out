package e

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"take-out/common"
)

var ErrMsg = map[int]string{
	SUCCESS:         "ok",
	ERROR:           "内部错误",
	UNKNOW_IDENTITY: "未知身份",
}

func GetMsg(code int) string {
	return ErrMsg[code]
}

func Send(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusOK, common.Result{Code: code, Msg: GetMsg(code)})
}
