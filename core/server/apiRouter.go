package server

import (
	"gin-init/core/wire"
)

func addApiRouter() {
	appController, _ := wire.InitializeApp()

	userGroup := apiGroup.Group("api")
	{
		userGroup.POST("info/create", appController.ApiController.Create)
		userGroup.POST("action/import", appController.ApiController.Import)
		// userGroup.POST("info/all", appController.ApiController.GetAllUser)
		// userGroup.GET("info/detail", appController.ApiController.GetUserDetail)
		// userGroup.POST("info/list", appController.ApiController.GetUserList)
		// userGroup.POST("info/update", appController.ApiController.UpdateUser)
		// userGroup.POST("info/update/passwd", appController.ApiController.UpdateUserPassword)
		// userGroup.POST("info/reset/passwd", appController.ApiController.ResetUserPassword)
		// userGroup.POST("info/delete", appController.ApiController.DeleteUser)
	}
}
