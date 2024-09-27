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
	"gin-init/config"
	"gin-init/model/entity"
	"gin-init/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func init() {
	// 连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	var err error
	DB, err = gorm.Open(mysql.Open(config.Conf.Mysql.GetDsn()), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println("[api init error]连接Mysql数据库失败, error=" + err.Error())
		return
	}
	// 连接成功
	fmt.Println("[Success] Mysql数据库连接成功！！！")
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
func main() {
	// TODO 确定 filePath 的来源方式
	// filePath := "D:/overall/project/api-mock/public/access.log"
	filePath := flag.String("file", "", "Path to the log file")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Please provide a log file path using the -file flag.")
		os.Exit(1)
	}

	logEntries, err := service.ReadAndParseLogFile(*filePath)
	if err != nil {
		return
	}

	importToDb(logEntries)
}
