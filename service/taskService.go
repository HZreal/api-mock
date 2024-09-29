package service

import (
	"gin-init/model/dto"
	"gin-init/model/entity"
	"gin-init/model/vo"
	"github.com/gin-gonic/gin"
	"log"
)

type TaskService struct {
	TaskModel *entity.TaskModel
}

func NewTaskService(taskModel *entity.TaskModel) *TaskService {
	return &TaskService{TaskModel: taskModel}
}

// Create
func (sv *TaskService) Create(c *gin.Context, body dto.TaskCreateDTO) vo.TaskDetailInfo {
	entityObj := entity.TaskModel{
		Name:       body.Name,
		TotalCount: body.TotalCount,
	}

	if result := db.Create(&entityObj); result.Error != nil {
		log.Printf("Failed to create entityObj, error: %v", result.Error)
		panic("failed to create entityObj")
	}
	return vo.TaskDetailInfo{
		Id:         entityObj.Id,
		Name:       entityObj.Name,
		TotalCount: entityObj.TotalCount,
	}
}

// GetList
func (sv *TaskService) GetList(c *gin.Context, query dto.QueryPagination, body map[string]interface{}) (result *vo.PaginationResult) {

	var entities []vo.TaskDetailInfo
	var total int64
	page, pageSize := query.Page, query.PageSize

	//
	offset := (page - 1) * pageSize

	// 获取数据总数和分页数据
	db.Model(sv.TaskModel).Where(body).Count(&total).Offset(offset).Limit(pageSize).Find(&entities)

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
		Records:     entities,
	}
}

// /////////////////////////////////////////// 以下接口待调整 //////////////////////////////////////////////////

func (sv *TaskService) GetAll(c *gin.Context, body dto.TaskCreateDTO) []vo.TaskDetailInfo {
	var entities []vo.TaskDetailInfo
	if err := db.Model(sv.TaskModel).Where(body).Find(&entities).Error; err != nil {
		log.Printf("query entities err:%v", err)
		panic(err)
	}
	return entities
}

func (sv *TaskService) GetDetail(c *gin.Context, id int) (userInfo vo.UserDetailInfo) {
	//
	affected := db.Take(sv.TaskModel, id).Scan(&userInfo).RowsAffected
	if affected == 0 {
		log.Printf("No entity found with ID: %d", id)
		return
	}

	return
}

//
// func (sv *TaskService) UpdateUser(c *gin.Context, body dto.UserUpdateDTO) vo.UserDetailInfo {
// 	id := body.Id
// 	var user entity.TaskModel
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
// func (sv *TaskService) DeleteUser(c *gin.Context, id int) {
// 	if result := db.First(sv.TaskModel, id); result.Error != nil {
// 		log.Printf("Failed to find user, error: %v", result.Error)
// 		panic("failed to find user")
// 	}
//
// 	//
// 	result := db.Delete(sv.TaskModel, id)
// 	if result.Error != nil {
// 		log.Printf("Failed to delete user, error: %v", result.Error)
// 		panic("failed to delete user")
// 	}
//
// }
