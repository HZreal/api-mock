package utils

/**
 * @Author elastic·H
 * @Date 2024-08-09
 * @File: generateId.go
 * @Description:
 */

import (
	"github.com/google/uuid"
)

func generateUUID() string {
	return uuid.New().String()
}
