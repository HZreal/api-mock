package utils

/**
 * @Author elastic·H
 * @Date 2024-08-09
 * @File: keyPair.go
 * @Description:
 */

import (
	"crypto/rand"
	"encoding/hex"
)

func generateKeyPair() (string, string, error) {
	// 生成 16 字节的随机数作为 accessKey
	accessKeyBytes := make([]byte, 16)
	_, err := rand.Read(accessKeyBytes)
	if err != nil {
		return "", "", err
	}
	accessKey := hex.EncodeToString(accessKeyBytes)

	// 生成 32 字节的随机数作为 secretKey
	secretKeyBytes := make([]byte, 32)
	_, err = rand.Read(secretKeyBytes)
	if err != nil {
		return "", "", err
	}
	secretKey := hex.EncodeToString(secretKeyBytes)

	return accessKey, secretKey, nil
}
