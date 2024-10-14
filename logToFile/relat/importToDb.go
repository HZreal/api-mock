package relat

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
	"github.com/samber/lo"
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
			existingApi.BodyType = line.BodyType
			existingApi.RequestBody = line.RequestBody

			for _, paramItem := range line.RequestBodyParams {
				_, existed := lo.Find(existingApi.Params, func(item *entity.ParamStruct) bool {
					return item.Name == paramItem.Name
				})
				if !existed {
					existingApi.Params = append(existingApi.Params, paramItem)
				}
			}

			if result := DB.Save(&existingApi); result.Error != nil {
				log.Printf("Failed to update api, error: %v", result.Error)
				continue
			}
		}

	}
}

// ParseAndImport 解析并导入
func ParseAndImport(filePath string) {
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
	ParseAndImport(filePath)
}
