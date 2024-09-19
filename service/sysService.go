package service

import (
	"errors"
	"gin-init/middleware"
	"gin-init/model/dto"
	"gin-init/model/vo"
	"github.com/gin-gonic/gin"
)

type SysService struct {
	UserService *UserService
}

func NewSysService(userService *UserService) *SysService {
	return &SysService{UserService: userService}
}

func (uS *SysService) Login(c *gin.Context, loginData dto.LoginData) (result vo.SysLoginVo, err error) {

	// 校验用户名密码
	// if !(loginData.Username == "admin" && loginData.Password == "root123456") {
	// 	return result, errors.New("invalid credentials")
	// }

	// TODO
	if !uS.UserService.CheckUser(loginData) {
		return result, errors.New("invalid credentials")
	}

	token, err := middleware.GenerateToken(loginData.Username)
	if err != nil {
		return result, errors.New("failed to generate token")
	}

	result = vo.SysLoginVo{Token: token}
	return result, nil
}

func (uS *SysService) Logout(c *gin.Context, id string) (data interface{}, err error) {
	return "", nil

}
