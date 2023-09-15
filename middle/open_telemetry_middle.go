package middle

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"take-out/common/utils"
)

func TelemetryMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer, closer := utils.NewJaeger("Gin-Scraper-Middle")
		engin := "rookie in jaeger"
		defer closer.Close()
		// 创建一个span并且设置tag
		span := tracer.StartSpan("middle")
		span.SetTag("engin", engin)
		// 设置启示Context
		opentracing.ContextWithSpan(c, span)
		c.Next()
	}
}
