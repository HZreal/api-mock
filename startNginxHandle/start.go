package main

/**
 * @Author elastic·H
 * @Date 2024-09-27
 * @File: start.go
 * @Description:
 */

import (
	"errors"
	"fmt"
	"gin-init/database"
	"gin-init/model/entity"
	"gin-init/service"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

func init() {
	DB = database.DB
}

// importToDb
func importToDb(logEntries []*service.Record) {
	for i, line := range logEntries {
		fmt.Println("i  ---->  ", i)

		var existingApi entity.ApiModel
		if err := DB.Where("name = ?", line.Name).First(&existingApi).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 记录不存在，创建新记录
				r := entity.ApiModel{
					Name:                line.Name,
					Method:              line.Method,
					UriArgs:             line.ReqUriArgs,
					Uri:                 line.Uri,
					ContentType:         line.ContentType,
					ResponseContentType: line.SentContentType,
					Args:                line.Args,
					BodyType:            line.BodyType,
					RequestBody:         line.RequestBody,
					Params:              line.RequestBodyParams,
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
			// TODO 记录已存在，执行更新操作，合并参数
			// existingApi.Name = line.ReqUriArgs
			// existingApi.Method = line.Method
			// existingApi.Uri = line.Uri
			// existingApi.UriArgs = line.ReqUriArgs
			existingApi.Args = line.Args
			existingApi.ContentType = line.ContentType
			existingApi.ResponseContentType = line.SentContentType
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

// func importToDb2(logEntries []*service.Record) {
// 	for i, line := range logEntries {
// 		fmt.Println("i  ---->  ", i)
//
// 		// TODO 改为 redis 解决重复性问题
// 		// 使用 SADD 添加元素并检查返回值
// 		ctx := context.Background()
// 		coding := line.Method + "_" + line.ReqUriArgs
//
// 		var existingApi entity.ApiModel
//
// 		result, err := rdb.SAdd(ctx, "okcc_api_set", coding).Result()
// 		if err != nil {
// 			log.Printf("Failed to SADD: %v", err)
// 			continue
// 		}
//
// 		if result == 1 {
// 			fmt.Println("Element was not present, successfully added.")
// 			// 记录不存在，创建新记录
// 			r := entity.ApiModel{
// 				Name:                coding,
// 				Method:              line.Method,
// 				UriArgs:             line.ReqUriArgs,
// 				Uri:                 line.Uri,
// 				ContentType:         line.ContentType,
// 				ResponseContentType: line.SentContentType,
// 				Args:                line.Args,
// 				BodyType:            line.BodyType,
// 				RequestBody:         line.RequestBody,
// 				Params:              line.RequestBodyParams,
// 			}
// 			if result2 := DB.Create(&r); result2.Error != nil {
// 				log.Printf("Failed to create api, error: %v", result2.Error)
// 				continue
// 			}
// 		} else if result == 0 {
// 			fmt.Println("Element already exists, no need to add.")
// 			if err := DB.Where("name = ?", coding).First(&existingApi).Error; err != nil {
// 				log.Printf("Failed to query api, error: %v", err)
// 				continue
// 			}
//
// 			// 记录已存在，执行更新操作
// 			// existingApi.Name = line.ReqUriArgs
// 			// existingApi.Method = line.Method
// 			// existingApi.UriArgs = line.ReqUriArgs
// 			// existingApi.Uri = line.Uri
// 			existingApi.ContentType = line.ContentType
// 			existingApi.Args = line.Args
// 			existingApi.BodyType = line.BodyType
// 			existingApi.RequestBody = line.RequestBody
// 			existingApi.Params = line.RequestBodyParams
//
// 			if result2 := DB.Save(&existingApi); result2.Error != nil {
// 				log.Printf("Failed to update api, error: %v", result2.Error)
// 				continue
// 			}
// 		}
//
// 	}
// }

// parseAndImport)
func parseAndImport(filePath string) {
	logEntries, err := service.ReadAndParseLogFile(filePath)
	if err != nil {
		return
	}

	importToDb(logEntries)
}

func main() {
	filePath := "D:/overall/project/api-mock/public/access.0930.log"
	// filePath := flag.String("file", "", "Path to the log file")
	// flag.Parse()
	//
	// if *filePath == "" {
	// 	fmt.Println("Please provide a log file path using the -file flag.")
	// 	os.Exit(1)
	// }
	// fmt.Println("filePath  ---->  ", *filePath)

	//
	parseAndImport(filePath)
}
