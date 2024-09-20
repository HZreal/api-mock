package service

/**
 * @Author elastic·H
 * @Date 2024-09-19
 * @File: nginxLogService.go
 * @Description:
 */

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type logLine struct {
	TimeLocal       string `json:"time_local"`
	Method          string `json:"request_method"`
	UriArgs         string `json:"request_uri"`
	Uri             string `json:"uri"`
	ContentType     string `json:"content_type"`
	Args            string `json:"args"`
	RequestBody     string `json:"request_body"`
	Status          string `json:"status"`
	RequestTime     string `json:"request_time"`
	SentContentType string `json:"sent_content_type"`
}

func getFilePath() string {
	return "D:/overall/project/api-mock/public/access.log"
}

func readAndParseLogFile() ([]logLine, error) {
	filePath := getFilePath()

	var logEntries []logLine

	//
	openFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer openFile.Close()

	//
	reader := bufio.NewReader(openFile)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		//
		var entry logLine
		if err := json.Unmarshal(line, &entry); err != nil {
			fmt.Printf("Unmarshal 失败: %s, 错误: %v\n", line, err)
			continue
		}

		// var params map[string]interface{}
		// err2 := json.Unmarshal([]byte(entry.RequestBody), &params)
		// if err2 != nil {
		// 	fmt.Printf("Unmarshal2 失败: %s, 错误: %v\n", entry.RequestBody, err)
		// 	continue
		// }

		logEntries = append(logEntries, entry)
	}
	return logEntries, nil
}
