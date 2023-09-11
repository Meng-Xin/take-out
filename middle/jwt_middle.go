package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"take-out/common"
	"take-out/common/e"
	"take-out/common/enum"
	"take-out/common/utils"
	"take-out/global"
)

func VerifyJWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		token := c.Request.Header.Get(global.Config.Jwt.Admin.Name)
		// 解析获取用户载荷信息
		payLoad, err := utils.VerifyToken(token)
		if err != nil {
			code = e.UNKNOW_IDENTITY
			c.JSON(http.StatusUnauthorized, common.Result{Code: code})
			c.Abort()
			return
		}
		// 在上下文设置载荷信息
		c.Set(enum.CurrentId, payLoad.UserID)
		c.Set(enum.CurrentName, payLoad.Username)
		// 这里是否要通知客户端重新保存新的Token
		c.Next()
	}
}

func VerifyJWTUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		token := c.Request.Header.Get(global.Config.Jwt.User.Name)
		// 解析获取用户载荷信息
		payLoad, err := utils.VerifyToken(token)
		if err != nil {
			code = e.UNKNOW_IDENTITY
			c.JSON(http.StatusUnauthorized, common.Result{Code: code})
			c.Abort()
			return
		}
		// 在上下文设置载荷信息
		c.Set(enum.CurrentId, payLoad.UserID)
		c.Set(enum.CurrentName, payLoad.Username)
		// 这里是否要通知客户端重新保存新的Token
		c.Next()
	}
}
