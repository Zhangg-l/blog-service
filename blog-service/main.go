package main

import (
	"fmt"
	"go_code/project8/blog-service/global"
	"go_code/project8/blog-service/internal/model"
	"go_code/project8/blog-service/internal/routers"
	"go_code/project8/blog-service/pkg/logger"
	"go_code/project8/blog-service/pkg/setting"
	"go_code/project8/blog-service/pkg/tracer"
	"log"
	"net/http"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	fmt.Println("---------------> let`go <---------------")
	db, err := model.NewDBEngine(global.DatabaseSetting)
	s := &http.Server{
		Addr:    ":8888",
		Handler: routers.NewRouter(),
	}

	s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Println(" init.setupSetting err:", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Println(" init.setupDBEngine() err:", err)
	}
	err = setupTracer()
	if err != nil {
		log.Println(" init.setupTracer() err:", err)
	}
	setupLogger()
}

// 给全局变量赋值
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.Readtimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.AppSetting.DefaultContextTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

// db启动
func setupDBEngine() error {
	db, err := model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.DBEngine = db
	return nil
}

// 日志启动
func setupLogger() {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath + "/" +
			global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

}

// db启动
func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "127.0.0.1:6831")

	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
