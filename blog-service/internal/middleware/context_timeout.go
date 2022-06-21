package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancle := context.WithTimeout(ctx.Request.Context(), t)
		defer cancle()
		// 把新建的context传下去
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	}
}
