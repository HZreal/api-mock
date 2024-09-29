package main

/**
 * @Author elastic·H
 * @Date 2024-09-27
 * @File: start.go
 * @Description:
 */

import (
	"errors"
	"flag"
	"fmt"
	"gin-init/database"
	"gin-init/model/entity"
	"gin-init/service"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func init() {
	DB = database.DB
}

// importToDb
func importToDb(logEntries []*service.Record) {
	for i, line := range logEntries {
		fmt.Println("i  ---->  ", i)

		//
		var existingApi entity.ApiModel
		if err := DB.Where("uri_args = ? AND method = ?", line.ReqUriArgs, line.Method).First(&existingApi).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 记录不存在，创建新记录
				r := entity.ApiModel{
					Name:        line.ReqUriArgs,
					Method:      line.Method,
					UriArgs:     line.ReqUriArgs,
					Uri:         line.Uri,
					ContentType: line.ContentType,
					Args:        line.Args,
					BodyType:    line.BodyType,
					RequestBody: line.RequestBody,
					Params:      line.RequestBodyParams,
				}
				if result := DB.Create(&r); result.Error != nil {
					log.Printf("Failed to create api, error: %v", result.Error)
					continue
				}
			} else {
				// 其他查询错误
				log.Printf("Failed to query api, error: %v", err)
				continue
			}
		} else {
			// 记录已存在，执行更新操作
			existingApi.Name = line.ReqUriArgs
			existingApi.Method = line.Method
			existingApi.Uri = line.Uri
			existingApi.ContentType = line.ContentType
			existingApi.Args = line.Args
			existingApi.BodyType = line.BodyType
			existingApi.RequestBody = line.RequestBody
			existingApi.Params = line.RequestBodyParams

			if result := DB.Save(&existingApi); result.Error != nil {
				log.Printf("Failed to update api, error: %v", result.Error)
				continue
			}
		}

		// 逐一入库，后续会更新
		// r := entity.ApiModel{
		// 	Name:        "xxx",
		// 	Method:      line.Method,
		// 	UriArgs:     line.ReqUriArgs,
		// 	Uri:         line.Uri,
		// 	ContentType: line.ContentType,
		// 	Args:        line.Args,
		// 	BodyType:    line.BodyType,
		// 	RequestBody: line.RequestBody,
		// 	Params:      line.RequestBodyParams,
		// }
		// if result := DB.Create(&r); result.Error != nil {
		// 	log.Printf("Failed to create api, error: %v", result.Error)
		// 	continue
		// }

	}
}

// parseAndImport)
func parseAndImport(filePath string) {
	logEntries, err := service.ReadAndParseLogFile(filePath)
	if err != nil {
		return
	}

	importToDb(logEntries)
}

func main() {
	// filePath := "D:/overall/project/api-mock/public/access.log"
	filePath := flag.String("file", "", "Path to the log file")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide a log file path using the -file flag.")
		os.Exit(1)
	}
	fmt.Println("filePath  ---->  ", *filePath)

	//
	parseAndImport(*filePath)
}
