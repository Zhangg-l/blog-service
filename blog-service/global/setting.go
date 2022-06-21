package global

import (
	"go_code/project8/blog-service/pkg/logger"
	"go_code/project8/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	JWTSetting      *setting.JWTSetting
	EmailSetting    *setting.EmailSetting
	Logger          *logger.Logger
)
