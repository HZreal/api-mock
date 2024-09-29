package database

import (
	"fmt"
	"gin-init/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresql() {
	// 连接 Postgresql 获取 db 实例
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN: config.Conf.Postgresql.GetDsn(), // data source name
		// PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	fmt.Println("[Success]  Postgresql 数据库连接成功！！！")

	// 设置数据库连接池参数
	sqlDB, _ := conn.DB()
	// 设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	sqlDB.SetMaxIdleConns(20)

	//
	DB = conn
}
