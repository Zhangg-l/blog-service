package v1

import (
	"go_code/project8/blog-service/global"
	"go_code/project8/blog-service/internal/service"
	"go_code/project8/blog-service/pkg/app"
	"go_code/project8/blog-service/pkg/convert"
	"go_code/project8/blog-service/pkg/errcode"
	"go_code/project8/blog-service/pkg/upload"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		response.ToErrorResponse(errcode.InValidParams.WithDetails(err.Error()))
		return
	}

	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InValidParams)
		return
	}
	svc := service.New(c.Request.Context())
	fileinfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)

	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile fail:%v", err)
		response.ToErrorResponse(errcode.Error_Upload_File_Fail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"file_access_url": fileinfo.AccessUrl,
	})

}
