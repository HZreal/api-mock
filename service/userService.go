package service

import (
	"encoding/json"
	"fmt"
	"gin-init/database"
	"gin-init/model/dto"
	"gin-init/model/entity"
	"gin-init/model/vo"
	"gin-init/service/common"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type UserService struct {
	UserModel    *entity.UserModel
	RedisService *common.RedisService // TODO 封装
}

func NewUserService(userModel *entity.UserModel, redisService *common.RedisService) *UserService {
	return &UserService{
		UserModel:    userModel,
		RedisService: redisService,
	}
}

func (uS *UserService) GetAllUser(c *gin.Context, body dto.UsersFilterDTO) []vo.UserDetailInfo {
	var users []vo.UserDetailInfo
	if err := database.DB.Model(uS.UserModel).Where(body).Find(&users).Error; err != nil {
		log.Printf("query users err:%v", err)
		panic(err)
	}
	return users
}

func (uS *UserService) GetUserList(c *gin.Context, query dto.QueryPagination, body map[string]interface{}) (result *vo.PaginationResult) {
	//

	var userInfos []vo.UserDetailInfo
	var total int64
	page, pageSize := query.Page, query.PageSize

	//
	offset := (page - 1) * pageSize

	// 获取数据总数和分页数据
	// database.DB.Model(&entity.UserModel{}).Where(body).Count(&total).Offset(offset).Limit(pageSize).Find(&userInfos)
	// TODO 通过依赖注入
	database.DB.Model(uS.UserModel).Where(body).Count(&total).Offset(offset).Limit(pageSize).Find(&userInfos)

	// 计算总页数
	pages := int(total) / pageSize
	if int(total)%pageSize != 0 {
		pages++
	}

	//
	return &vo.PaginationResult{
		Total:       int(total),
		Pages:       pages,
		CurrentPage: page,
		PageSize:    pageSize,
		Records:     userInfos,
	}
}

func (uS *UserService) GetUserDetail(c *gin.Context, id int) (userInfo vo.UserDetailInfo) {
	//
	key := fmt.Sprintf("tmp:user:id:%d", id)
	cachedData, err := uS.RedisService.Client.Get(c, key).Result()
	if err == redis.Nil {
		// 无缓存
		affected := database.DB.Take(&entity.UserModel{}, id).Scan(&userInfo).RowsAffected
		if affected == 0 {
			log.Printf("No user found with ID: %d", id)
			return
		}

		// 将查询结果序列化为 JSON 字符串
		unitInfoJson, err := json.Marshal(userInfo)
		if err != nil {
			log.Printf("Failed to serialize data for user ID: %d, error: %v", id, err)
			panic("failed to serialize data")
		}

		// 将数据缓存到 Redis，设置缓存过期时间为 30 S
		err = uS.RedisService.Client.Set(c, key, unitInfoJson, 30*time.Second).Err()
		if err != nil {
			panic("failed to save data")
		}

		return userInfo

	} else if err != nil {
		log.Printf("Failed to get cache for key: %s, error: %v", key, err)
		panic("failed to get cache")
	} else {
		// 如果缓存中有数据，返回缓存数据
		if err := json.Unmarshal([]byte(cachedData), &userInfo); err != nil {
			log.Printf("Failed to deserialize cache data for user ID: %d, error: %v", id, err)
			panic("failed to deserialize cache data")
		}

		return userInfo
	}

}

func (uS *UserService) CheckUser(loginData dto.LoginData) bool {
	//
	username, password := loginData.Username, loginData.Password

	// TODO（这里应该是从数据库中验证用户信息）
	if !(username == "admin" && password == "root123456") {
		return false
	}
	return true
}

func (uS *UserService) CreateUser(c *gin.Context, body dto.UserCreateDTO) vo.UserDetailInfo {
	user := entity.UserModel{
		Username: body.Username,
		Password: body.Password,
		Phone:    body.Phone,
		Age:      body.Age,
	}

	if result := database.DB.Create(&user); result.Error != nil {
		log.Printf("Failed to create user, error: %v", result.Error)
		panic("failed to create user")
	}
	return vo.UserDetailInfo{
		Id:       user.Id,
		Username: user.Username,
		Phone:    user.Phone,
		Age:      user.Age,
	}

}

func (uS *UserService) UpdateUser(c *gin.Context, body dto.UserUpdateDTO) vo.UserDetailInfo {
	id := body.Id
	var user entity.UserModel
	if result := database.DB.First(&user, id); result.Error != nil {
		log.Printf("Failed to find user, error: %v", result.Error)
		panic("failed to find user")
	}

	//
	result := database.DB.Model(&user).Where("id = ?", id).Updates(body)
	if result.Error != nil {
		log.Printf("Failed to update user, error: %v", result.Error)
		panic("failed to update user")
	}
	return vo.UserDetailInfo{
		Id:       user.Id,
		Username: user.Username,
		Phone:    user.Phone,
		Age:      user.Age,
	}
}

func (uS *UserService) DeleteUser(c *gin.Context, id int) {
	if result := database.DB.First(uS.UserModel, id); result.Error != nil {
		log.Printf("Failed to find user, error: %v", result.Error)
		panic("failed to find user")
	}

	//
	result := database.DB.Delete(uS.UserModel, id)
	if result.Error != nil {
		log.Printf("Failed to delete user, error: %v", result.Error)
		panic("failed to delete user")
	}

}
