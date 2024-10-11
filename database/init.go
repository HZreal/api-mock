package database

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client

// var Ctx = context.Background()

func init() {
	// InitRedis()
	InitMysql()
}
