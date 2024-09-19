package server

import (
	"gin-init/core/wire"
)

func addDemoRouter() {
	// demoController := controller.DemoController{}
	appController, _ := wire.InitializeApp()

	sysGroup := apiGroup.Group("demo")
	{
		sysGroup.GET("sendMQ", appController.DemoController.SendMsgWithRabbitMQ)
	}

}
