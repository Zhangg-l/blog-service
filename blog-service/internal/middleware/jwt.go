package middleware

import (
	"fmt"
	"go_code/project8/blog-service/pkg/app"
	"go_code/project8/blog-service/pkg/errcode"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			ecode = errcode.Success
			token string
		)

		if s, exit := ctx.GetQuery("token"); exit {
			token = s
		} else {
			token = ctx.GetHeader("token")
		}
		fmt.Println("token:",token)
		fmt.Println("========")
		if token == "" {
			ecode = errcode.InValidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(ctx)
			response.ToErrorResponse(ecode)
			ctx.Abort()
			return
		}

		ctx.Next()
	}

}
