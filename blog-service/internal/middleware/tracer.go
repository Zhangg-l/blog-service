package middleware

import (
	"context"
	"go_code/project8/blog-service/global"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx context.Context
		var traceId string
		var spanId string
		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			span, ctx =
				opentracing.StartSpanFromContextWithTracer(c.Request.Context(),
					global.Tracer, c.Request.URL.Path,
					opentracing.ChildOf(span.Context()))

		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(),
				global.Tracer, c.Request.URL.Path)
		}
		defer span.Finish()

		spanCtx := span.Context()

		switch spanCtx.(type) {
		case jaeger.SpanContext:
			traceId = spanCtx.(jaeger.SpanContext).TraceID().String()
			spanId = spanCtx.(jaeger.SpanContext).SpanID().String()
		}
		c.Set("X-Trace-ID", traceId)
		c.Set("X-Span-Id", spanId)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
