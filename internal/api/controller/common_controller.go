package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"take-out/common"
	"take-out/common/e"
	"take-out/common/utils"
	"take-out/global"
)

type CommonController struct {
}

func (cc *CommonController) Upload(ctx *gin.Context) {
	code := e.SUCCESS
	// 获取前端传递的图片
	file, err := ctx.FormFile("file")
	if err != nil {
		return
	}
	// 拼接uuid的图片名称
	uuid := uuid.New()
	imageName := uuid.String() + file.Filename
	imagePath, err := utils.AliyunOss(imageName, file)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("AliyunOss upload failed", "err", err.Error())
	}
	ctx.JSON(http.StatusOK, common.Result{Code: code, Data: imagePath, Msg: e.GetMsg(code)})
}
