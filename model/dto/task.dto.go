package dto

/**
 * @Author elastic·H
 * @Date 2024-09-25
 * @File: task.dto.go
 * @Description:
 */

type TaskCreateDTO struct {
	Name       string `json:"name" binding:"omitempty,min=1,max=20"`
	TotalCount uint32 `json:"totalCount" binding:"required,min=0"`
}
