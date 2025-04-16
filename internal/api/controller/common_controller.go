package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"take-out/common/retcode"
	"take-out/common/utils"
	"take-out/global"
)

type CommonController struct {
}

func (cc *CommonController) Upload(ctx *gin.Context) {
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
		global.Log.Warn("AliyunOss upload failed", "err", err.Error())
		retcode.Fatal(ctx, err, "")
		return
	}
	retcode.OK(ctx, imagePath)
}
