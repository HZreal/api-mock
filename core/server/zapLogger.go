package server

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

/**
 * @Author elastic·H
 * @Date 2024-09-08
 * @File: zapLogger.go
 * @Description:
 */

var Logger *zap.Logger

// func init2() {
// 	// 开发环境下使用 zap.NewDevelopment()，生产环境下使用 zap.NewProduction()
// 	Logger, _ = zap.NewDevelopment()
// 	// Logger, _ = zap.NewProduction()
// }

func init() {
	logFile, err := os.OpenFile("logs/gin_app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	// 创建自定义的 zapcore，输出到文件
	writeSyncer := zapcore.AddSync(logFile)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 输出 JSON 格式的日志
		writeSyncer,                           // 输出到日志文件
		zapcore.InfoLevel,                     // 日志级别
	)

	Logger = zap.New(core, zap.AddCaller())
}
