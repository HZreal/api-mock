package routes

import (
	"gin-init/core/wire"
	"github.com/gin-gonic/gin"
)

func AddDemoRouter(r *gin.RouterGroup) {
	// demoController := controller.DemoController{}
	appController, _ := wire.InitializeApp()

	sysGroup := r.Group("demo")
	{
		sysGroup.GET("sendMQ", appController.DemoController.SendMsgWithRabbitMQ)
	}

}
