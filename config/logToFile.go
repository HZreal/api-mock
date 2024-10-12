package config

/**
 * @Author elastic·H
 * @Date 2024-10-12
 * @File: logToFile.go
 * @Description:
 */

type LogToFileConfig struct {
	LogSourcePath      string `yaml:"logSourcePath"`
	CollectionStoreDir string `yaml:"collectionStoreDir"`
}
