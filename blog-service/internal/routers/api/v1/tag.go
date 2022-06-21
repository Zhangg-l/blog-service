package v1

import (
	"fmt"
	"go_code/project8/blog-service/global"
	"go_code/project8/blog-service/internal/service"
	"go_code/project8/blog-service/pkg/app"
	"go_code/project8/blog-service/pkg/convert"
	"go_code/project8/blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}
func (t Tag) List(c *gin.Context) {

	// 定义入参
	param := service.TagListRequest{}
	// 定义响应
	response := app.NewResponse(c)
	// 入参和绑定的校验
	valid, errs := app.BindAndValid(c, &param)
	// 参数校验失败
	if valid {
		global.Logger.Errorf(c, "app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	// 获取页面信息
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})

	if err != nil {
		global.Logger.Errorf(c, "svc.CountTag errs:%v", errs)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList errs:%v", errs)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
	return
}

func (t Tag) Create(c *gin.Context) {
	fmt.Println(c == c.Request.Context())
	param := service.CreateTagRequest{}
	resp := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	// 参数校验失败
	if valid {
		global.Logger.Errorf(c, "app.BindAndValid errs:%v", errs)
		resp.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateTag err:%v", errs)
		resp.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	resp.ToResponse(gin.H{})
	return
}
func (t Tag) Update(c *gin.Context) {

	param := service.UpdateTagRequest{
		Id: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	resp := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	// 参数校验失败
	if valid {
		global.Logger.Errorf(c, "app.BindAndValid errs:%v", errs)
		resp.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateTag err:%v", errs)
		resp.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	resp.ToResponse(gin.H{})
	return
}
func (t Tag) Delete(c *gin.Context) {

	param := service.DeleteTagRequest{
		Id: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	resp := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	// 参数校验失败
	if valid {
		global.Logger.Errorf(c, "app.BindAndValid errs:%v", errs)
		resp.ToErrorResponse(errcode.InValidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteTag err:%v", errs)
		resp.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	resp.ToResponse(gin.H{})
	return
}
