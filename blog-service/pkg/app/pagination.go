package app

import (
	"fmt"
	"go_code/project8/blog-service/global"
	"go_code/project8/blog-service/pkg/convert"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 0
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	fmt.Print(c.Query("page_size"))
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}
func GetPageOffset(page, pageSize int) int {
	res := 0
	if page > 0 {
		res = (page - 1) * pageSize
	}
	return res
}
