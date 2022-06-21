package routers

import (
	"go_code/project8/blog-service/global"
	"go_code/project8/blog-service/internal/middleware"
	v1 "go_code/project8/blog-service/internal/routers/api/v1"
	"go_code/project8/blog-service/pkg/limiter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var methodLimiter = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterfaceBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Revovery())
	}
	r.Use(middleware.Tracing())
	r.Use(middleware.RateLimiter(methodLimiter))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translation())
	r.POST("/upload/file", v1.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.GET("/auth", v1.GetAuth)
	tag := v1.Tag{}

	aticle := v1.Article{}

	apiV1 := r.Group("api/v1")
	apiV1.Use(middleware.JWT())
	{
		// 获取标签
		apiV1.GET("/tags", tag.List)
		// 增加标签
		apiV1.POST("/tags", tag.Create)
		// 删除指定标签
		apiV1.DELETE("/tags/:id", tag.Delete)
		// 更新指定标签
		apiV1.PUT("/tags/:id", tag.Update)
		// 更新一个资源的一部分
		apiV1.PATCH("/tags/:id/state", tag.Update)

		apiV1.GET("/articles", aticle.List)
		apiV1.POST("/articles", aticle.Create)
		apiV1.DELETE("/articles/:id", aticle.Delete)
		apiV1.PUT("/articles/:id", aticle.Update)
		apiV1.PATCH("/articles/:id/state", aticle.Update)
		apiV1.GET("/articles/:id", aticle.Get)

	}
	return r
}
