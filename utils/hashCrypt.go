package utils

/**
 * @Author elastic·H
 * @Date 2024-08-09
 * @File: hashCrypt.go
 * @Description:
 */

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 对密码进行哈希计算
func HashPassword(password string) (string, error) {
	// 生成哈希值，cost 值越高计算越耗时，默认为 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	fmt.Println("hashedPassword  ---->  ", string(hashedPassword))
	return string(hashedPassword), nil
}

// CheckPasswordHash 验证密码是否匹配
func CheckPasswordHash(password, hashedPassword string) bool {
	// 将用户输入的密码与数据库中存储的哈希密码进行比较
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func encryptAndValidate(pwd string) bool {
	password, _ := HashPassword("123456")
	return CheckPasswordHash("123456", password)
}
