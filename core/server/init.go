package server

import (
	"fmt"
	"gin-init/config"
	"gin-init/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

var r *gin.Engine

var apiGroup *gin.RouterGroup

func init() {
	//
	gin.SetMode(config.Conf.Gin.Mode)

	//
	r = gin.Default()

	// 日志
	// r.Use(gin.Logger())
	// 将 Zap 日志器作为 Gin 的日志中间件
	// r.Use(ginzap.Ginzap(Logger, time.RFC3339, true))
	// r.Use(ginzap.RecoveryWithZap(Logger, true))

	// 异常拦截
	// r.Use(gin.Recovery())
	r.Use(middleware.ExceptionInterceptorMiddleware())

	// cors
	// 全局注册 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 允许的域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 允许的HTTP方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // 允许携带认证信息
		MaxAge:           12 * time.Hour, // 预检请求的缓存时间
	}))

	// 静态文件（页面等）
	r.Static("/public", "./public")

	// routers definition
	apiGroup = r.Group("api")
	//
	registerRoutes()
}

func registerRoutes() {
	addSysRouter()
	addDemoRouter()
	addUserRouter()
	addApiRouter()
}

func StartGinServer() {
	err := r.Run(config.Conf.Gin.GetAddr())
	if err != nil {
		fmt.Println("[Error] r.Run " + err.Error())
		return
	}
}
