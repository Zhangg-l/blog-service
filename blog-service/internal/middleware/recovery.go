package middleware

import (
	"fmt"
	"go_code/project8/blog-service/global"
	email "go_code/project8/blog-service/pkg/Email"
	"go_code/project8/blog-service/pkg/app"
	"go_code/project8/blog-service/pkg/errcode"
	"time"

	"github.com/gin-gonic/gin"
)

func Revovery() gin.HandlerFunc {
	defailMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallerFrames().Errorf(ctx, "panic recover err:%v", err)
				err := defailMailer.SendMail(global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间：%d", time.Now().Unix()),
					fmt.Sprintf("错误信息:%v", err))
				if err != nil {
					global.Logger.Panicf(ctx, "mail.sendMail err :%v", err)
				}
				app.NewResponse(ctx).ToErrorResponse(errcode.ServeError)
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
