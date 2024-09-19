package server

import (
	"gin-init/core/wire"
)

func addSysRouter() {
	// sysController := controller.SysController{}
	appController, _ := wire.InitializeApp()

	sysGroup := apiGroup.Group("sys")
	{
		// sysGroup.POST("login", sysController.Login)
		sysGroup.POST("login", appController.SysController.Login)
		sysGroup.POST("logout", appController.SysController.Logout)
		sysGroup.POST("config")
		sysGroup.POST("config/set")
	}

}
