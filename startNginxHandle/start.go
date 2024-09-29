package main

/**
 * @Author elastic·H
 * @Date 2024-09-27
 * @File: start.go
 * @Description:
 */

import (
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

		// 逐一入库，后续会更新
		r := entity.ApiModel{
			Name:        "xxx",
			Method:      line.Method,
			UriArgs:     line.ReqUriArgs,
			Uri:         line.Uri,
			ContentType: line.ContentType,
			Args:        line.Args,
			Params:      line.RequestBodyParams,
		}
		if result := DB.Create(&r); result.Error != nil {
			log.Printf("Failed to create api, error: %v", result.Error)
			continue
		}

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
