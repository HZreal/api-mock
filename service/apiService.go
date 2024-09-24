package service

import (
	"encoding/json"
	"fmt"
	"gin-init/model/dto"
	"gin-init/model/entity"
	"gin-init/model/vo"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type ApiService struct {
	ApiModel *entity.ApiModel
}

func NewApiService(apiModel *entity.ApiModel) *ApiService {
	return &ApiService{ApiModel: apiModel}
}

func (uS *ApiService) Create(c *gin.Context, body dto.ApiCreateDTO) vo.ApiDetailInfo {
	api := entity.ApiModel{
		Name:   body.Name,
		Uri:    body.Url,
		Method: body.Method,
	}

	if result := db.Create(&api); result.Error != nil {
		log.Printf("Failed to create api, error: %v", result.Error)
		panic("failed to create api")
	}
	return vo.ApiDetailInfo{
		Id:     api.Id,
		Name:   api.Name,
		Method: api.Method,
		Params: api.Params,
	}
}

func (uS *ApiService) Import() {
	logEntries, err := readAndParseLogFile()
	if err != nil {
		return
	}

	for i, line := range logEntries {
		fmt.Println("i  ---->  ", i)

		//
		var body map[string]interface{}
		var params []entity.ParamStruct

		lineBody := line.RequestBody
		if lineBody == "" {
			//
			params = []entity.ParamStruct{}
		} else if strings.HasPrefix(lineBody, "p=") {
			//
			body = ParseBodyP(lineBody)
		} else if strings.Contains(lineBody, "=") {
			//
			body, _ = ParseURLFormEncoded(lineBody)
		} else {
			err2 := json.Unmarshal([]byte(lineBody), &body)
			if err2 != nil {
				fmt.Println("lineBody Unmarshal 失败, params 为空", err2)
				params = []entity.ParamStruct{}
			} else {
				//
				fmt.Println("lineBody Unmarshal成功 ---->  ", body)
			}
		}

		// TODO 处理 body 成参数
		for k, v := range body {
			var itemType string
			switch v.(type) {
			case string:
				itemType = "string"
			case int:
				itemType = "int"
			case float32:
				itemType = "float32"
			case float64:
				itemType = "float64"
			case bool:
				itemType = "bool"
			}
			item := entity.ParamStruct{Name: k, Type: itemType}
			params = append(params, item)
		}

		//
		r := entity.ApiModel{
			Name:        "xxx",
			Method:      line.Method,
			UriArgs:     line.ReqUriArgs,
			Uri:         line.Uri,
			ContentType: line.ContentType,
			Args:        line.Args,
			Params:      params,
		}
		if result := db.Create(&r); result.Error != nil {
			log.Printf("Failed to create api, error: %v", result.Error)
			continue
		}

	}

}

// /////////////////////////////////////////// 以下接口待调整 //////////////////////////////////////////////////

// func (uS *ApiService) GetAll(c *gin.Context, body dto.UsersFilterDTO) []vo.UserDetailInfo {
// 	var users []vo.UserDetailInfo
// 	if err := db.Model(uS.ApiModel).Where(body).Find(&users).Error; err != nil {
// 		log.Printf("query users err:%v", err)
// 		panic(err)
// 	}
// 	return users
// }
//
// func (uS *ApiService) GetUserList(c *gin.Context, query dto.QueryPagination, body map[string]interface{}) (result *vo.PaginationResult) {
// 	//
//
// 	var userInfos []vo.UserDetailInfo
// 	var total int64
// 	page, pageSize := query.Page, query.PageSize
//
// 	//
// 	offset := (page - 1) * pageSize
//
// 	// 获取数据总数和分页数据
// 	// db.Model(&entity.ApiModel{}).Where(body).Count(&total).Offset(offset).Limit(pageSize).Find(&userInfos)
// 	// TODO 通过依赖注入
// 	db.Model(uS.ApiModel).Where(body).Count(&total).Offset(offset).Limit(pageSize).Find(&userInfos)
//
// 	// 计算总页数
// 	pages := int(total) / pageSize
// 	if int(total)%pageSize != 0 {
// 		pages++
// 	}
//
// 	//
// 	return &vo.PaginationResult{
// 		Total:       int(total),
// 		Pages:       pages,
// 		CurrentPage: page,
// 		PageSize:    pageSize,
// 		Records:     userInfos,
// 	}
// }
//
// func (uS *ApiService) GetUserDetail(c *gin.Context, id int) (userInfo vo.UserDetailInfo) {
// 	//
// 	key := fmt.Sprintf("tmp:user:id:%s", id)
// 	cachedData, err := rdb.Get(c, key).Result()
// 	if err == redis.Nil {
// 		// 无缓存
// 		affected := db.Take(&entity.ApiModel{}, id).Scan(&userInfo).RowsAffected
// 		if affected == 0 {
// 			log.Printf("No user found with ID: %s", id)
// 			return
// 		}
//
// 		// 将查询结果序列化为 JSON 字符串
// 		unitInfoJson, err := json.Marshal(userInfo)
// 		if err != nil {
// 			log.Printf("Failed to serialize data for user ID: %s, error: %v", id, err)
// 			panic("failed to serialize data")
// 		}
//
// 		// 将数据缓存到 Redis，设置缓存过期时间为 30 S
// 		err = rdb.Set(c, key, unitInfoJson, 30*time.Second).Err()
// 		if err != nil {
// 			panic("failed to save data")
// 		}
//
// 		return userInfo
//
// 	} else if err != nil {
// 		log.Printf("Failed to get cache for key: %s, error: %v", key, err)
// 		panic("failed to get cache")
// 	} else {
// 		// 如果缓存中有数据，返回缓存数据
// 		if err := json.Unmarshal([]byte(cachedData), &userInfo); err != nil {
// 			log.Printf("Failed to deserialize cache data for user ID: %s, error: %v", id, err)
// 			panic("failed to deserialize cache data")
// 		}
//
// 		return userInfo
// 	}
//
// }
//
// func (uS *ApiService) UpdateUser(c *gin.Context, body dto.UserUpdateDTO) vo.UserDetailInfo {
// 	id := body.Id
// 	var user entity.ApiModel
// 	if result := db.First(&user, id); result.Error != nil {
// 		log.Printf("Failed to find user, error: %v", result.Error)
// 		panic("failed to find user")
// 	}
//
// 	//
// 	result := db.Model(&user).Where("id = ?", id).Updates(body)
// 	if result.Error != nil {
// 		log.Printf("Failed to update user, error: %v", result.Error)
// 		panic("failed to update user")
// 	}
// 	return vo.UserDetailInfo{
// 		Id:       user.Id,
// 		Username: user.Username,
// 		Phone:    user.Phone,
// 		Age:      user.Age,
// 	}
// }
//
// func (uS *ApiService) DeleteUser(c *gin.Context, id int) {
// 	if result := db.First(uS.ApiModel, id); result.Error != nil {
// 		log.Printf("Failed to find user, error: %v", result.Error)
// 		panic("failed to find user")
// 	}
//
// 	//
// 	result := db.Delete(uS.ApiModel, id)
// 	if result.Error != nil {
// 		log.Printf("Failed to delete user, error: %v", result.Error)
// 		panic("failed to delete user")
// 	}
//
// }
