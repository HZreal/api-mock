package server

import (
	"gin-init/core/wire"
)

func addApiRouter() {
	appController, _ := wire.InitializeApp()

	apiGroup := routerGroup.Group("api")
	{
		apiGroup.POST("info/create", appController.ApiController.Create)
		apiGroup.POST("action/import", appController.ApiController.Import)
		// apiGroup.POST("info/all", appController.ApiController.GetAllUser)
		// apiGroup.GET("info/detail", appController.ApiController.GetUserDetail)
		apiGroup.POST("info/list", appController.ApiController.GetList)
		// apiGroup.POST("info/update", appController.ApiController.UpdateUser)
		// apiGroup.POST("info/update/passwd", appController.ApiController.UpdateUserPassword)
		// apiGroup.POST("info/reset/passwd", appController.ApiController.ResetUserPassword)
		// apiGroup.POST("info/delete", appController.ApiController.DeleteUser)
	}
}
