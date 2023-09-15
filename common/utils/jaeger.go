package utils

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

func NewJaeger(service string) (opentracing.Tracer, io.Closer) {
	// trace 配置
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			// collector 信息根据自己ip配置
			CollectorEndpoint: "http://127.0.0.1:14268/api/traces",
		},
	}
	// 根据上面的配置新建一个tracer
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

func LinkNext(c context.Context, tag string, linkNode string) {
	span, _ := opentracing.StartSpanFromContext(c, "formatString")
	defer span.Finish()
	span.SetTag(tag, linkNode)
	span.LogFields(log.String(linkNode, "Next……"))
}
