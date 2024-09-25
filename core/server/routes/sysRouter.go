package routes

import (
	"gin-init/core/wire"
	"github.com/gin-gonic/gin"
)

func AddSysRouter(r *gin.RouterGroup) {
	// sysController := controller.SysController{}
	appController, _ := wire.InitializeApp()

	sysGroup := r.Group("sys")
	{
		// sysGroup.POST("login", sysController.Login)
		sysGroup.POST("login", appController.SysController.Login)
		sysGroup.POST("logout", appController.SysController.Logout)
		sysGroup.POST("config")
		sysGroup.POST("config/set")
	}

}
