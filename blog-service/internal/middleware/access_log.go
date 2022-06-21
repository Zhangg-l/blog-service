package middleware

import (
	"bytes"
	"go_code/project8/blog-service/global"
	"go_code/project8/blog-service/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// 实现 responsewriter
type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (a AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := a.body.Write(p); err != nil {
		return n, err
	}
	return a.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		// 做替换 让ctx传下去的是自己的实现，可以从自己的实现中取出想要的数据
		ctx.Writer = bodyWriter
		begineTime := time.Now().Unix()
		ctx.Next()
		endTime := time.Now().Unix()
		fields := logger.Fields{
			"request":  ctx.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof(ctx, "access log:method-->%s,status_code--> %d,begine_time-->%d,end_time-->%d", ctx.Request.Method,
			bodyWriter.Status(), begineTime, endTime)

	}
}
