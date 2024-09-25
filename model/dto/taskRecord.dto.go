package dto

/**
 * @Author elastic·H
 * @Date 2024-09-25
 * @File: task.dto.go
 * @Description:
 */

type TaskRecordCreateDTO struct {
	TaskId uint64 `json:"task_id" binding:"required,min=0"`
	ApiId  uint64 `json:"api_id" binding:"required,min=0"`
}
