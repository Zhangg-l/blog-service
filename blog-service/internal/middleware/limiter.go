package middleware

import (
	"go_code/project8/blog-service/pkg/app"
	"go_code/project8/blog-service/pkg/errcode"
	"go_code/project8/blog-service/pkg/limiter"

	"github.com/gin-gonic/gin"
)

// 限流器
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := l.Key(ctx)

		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				resp := app.NewResponse(ctx)
				resp.ToErrorResponse(errcode.TooManyRequests)
				ctx.Abort()
				return
			}
		}
		ctx.Next()

	}
}
