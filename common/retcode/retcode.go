package retcode

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"take-out/common/e"
)

// ErrorCodeGetter 错误接口：用于提取错误码
type ErrorCodeGetter interface {
	GetCode() int
}

// Error 通用错误结构体
type Error struct {
	ErrCode int
	ErrMsg  string
}

func (e *Error) Error() string {
	return e.ErrMsg
}

func (e *Error) GetCode() int {
	return e.ErrCode
}

func NewError(code int, msg string) *Error {
	return &Error{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

// OK 渲染成功响应
func OK(c *gin.Context, output interface{}) {
	RenderReply(c, output)
}

// Fatal 渲染服务错误响应
func Fatal(c *gin.Context, e error, msg string) {
	code := GetErrCode(e)
	fmt.Printf("%s %+v", msg, e)
	if msg == "" {
		msg = e.Error()
	}
	RenderErrMsg(c, code, msg)
}

// CustomError 渲染自定义错误
func CustomError(c *gin.Context, code int, msg string) {
	fmt.Printf("%d %s", code, msg)
	RenderErrMsg(c, code, msg)
}

// RenderReply
func RenderReply(c *gin.Context, data interface{}) {
	render(c, e.SUCCESS, data, nil)
}

// RenderErrMsg
func RenderErrMsg(c *gin.Context, code int, msg string) {
	render(c, code, nil, errors.New(msg))
}

// render
func render(c *gin.Context, code int, data interface{}, err error) {
	var msg string
	if err != nil {
		msg = err.Error()
		if code == e.UNKNOW_IDENTITY {
			code = GetErrCode(err)
		}
	} else if defaultMsg, ok := e.ErrMsg[code]; ok {
		msg = defaultMsg
	}

	r := gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	c.Set("return_code", code)
	c.JSON(http.StatusOK, r)
}

// GetErrCode 获取错误码
func GetErrCode(err error) int {
	if errGetter, ok := err.(ErrorCodeGetter); ok {
		return errGetter.GetCode()
	}

	switch errType := err.(type) {
	case *mysql.MySQLError:
		return int(errType.Number)
	case *Error:
		return errType.ErrCode
	}

	return e.UNKNOW_IDENTITY
}
