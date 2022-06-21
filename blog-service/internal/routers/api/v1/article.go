package v1

import (
	"go_code/project8/blog-service/global"
	"go_code/project8/blog-service/internal/service"
	"go_code/project8/blog-service/pkg/app"
	"go_code/project8/blog-service/pkg/convert"
	"go_code/project8/blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Id            uint32 `json:"id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	Desc          string `json:"desc"`
	State         uint8  `json:"state"`
}

func NewArticle() Article {
	return Article{}
}

func (t Article) Get(c *gin.Context) {

	param := service.ArticleRequest{
		Id: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		global.Logger.Errorf(c, "app.BindAndValid err :%v", errs)
		response.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticle err :%v", err)
		response.ToErrorResponse(errcode.ErrorGetAticleFail)
		return
	}
	response.ToResponse(article)
}
func (t Article) List(c *gin.Context) {

	param := service.ArticleListRequest{}
	resp := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if valid {
		global.Logger.Errorf(c, "app.BindAndValid err :%v", errs)
		resp.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	articleList, total, err := svc.GetArticleList(&param, &pager)

	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticleList err :%v", err)
		resp.ToErrorResponse(errcode.ErrorGetAticleListFail)
		return

	}
	resp.ToResponseList(articleList, total)

}
func (t Article) Create(c *gin.Context) {

	resp := app.NewResponse(c)

	param := service.CreateArticleRequest{}

	valid, errs := app.BindAndValid(c, &param)
	if valid {
		global.Logger.Errorf(c, "app.BindAndValid err :%v", errs)
		resp.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateArticle err :%v", err)
		resp.ToErrorResponse(errcode.ErrorCreateAticleFail)
		return
	}
	resp.ToResponse(gin.H{})
	return
}
func (t Article) Update(c *gin.Context) {

	resp := app.NewResponse(c)

	param := service.UpdateArticleRequest{
		Id: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	valid, errs := app.BindAndValid(c, &param)
	if valid {
		global.Logger.Errorf(c, "app.BindAndValid err :%v", errs)
		resp.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateArticle err :%v", err)
		resp.ToErrorResponse(errcode.ErrorUpdateAticleFail)
		return
	}
	resp.ToResponse(gin.H{})
	return
}
func (t Article) Delete(c *gin.Context) {

	param := service.DeleteArticleRequest{
		Id: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	resp := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if valid {
		global.Logger.Errorf(c, "app.BindAndValid err :%v", errs)
		resp.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteArticle err :%v", err)
		resp.ToErrorResponse(errcode.ErrorUpdateAticleFail)
		return
	}
	resp.ToResponse(gin.H{})
	return

}
