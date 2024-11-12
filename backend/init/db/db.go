package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/1Panel-dev/1Panel/backend/global"
	"gorm.io/driver/mysql" // 导入 MySQL 驱动
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() {
	fmt.Println(global.CONF.System.DbPath)
	// 初始化数据库路径，如果不需要，也可以去掉这部分
	if _, err := os.Stat(global.CONF.System.DbPath); err != nil {
		if err := os.MkdirAll(global.CONF.System.DbPath, os.ModePerm); err != nil {
			panic(fmt.Errorf("init db dir failed, err: %v", err))
		}
	}

	// MySQL 连接字符串，配置数据库的用户、密码、主机、端口和数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",      // MySQL 用户名
		"root",      // MySQL 密码
		"127.0.0.1", // MySQL 主机
		"3306",      // MySQL 端口
		"ZYC",       // 需要连接的数据库名，修改为你的数据库名称
	)

	// 配置 MySQL 日志输出
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	// 初始化监控数据库部分
	initMonitorDB(newLogger)

	// 使用 MySQL 驱动连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	})
	if err != nil {
		panic(fmt.Errorf("init db failed, err: %v", err))
	}

	// 配置数据库连接池
	sqlDB, dbError := db.DB()
	if dbError != nil {
		panic(dbError)
	}
	sqlDB.SetConnMaxIdleTime(10 * time.Second)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 将数据库对象赋值到 global.DB 中
	global.DB = db
	global.LOG.Info("init db successfully")
}

func initMonitorDB(newLogger logger.Interface) {
	// 使用 MySQL 数据源替换原来的 SQLite
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",      // MySQL 用户名
		"root",      // MySQL 密码
		"127.0.0.1", // MySQL 主机
		"3306",      // MySQL 端口
		"ZYC",       // 用于监控的数据库名，修改为你的数据库名称
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	})
	if err != nil {
		panic(fmt.Errorf("init monitor db failed, err: %v", err))
	}

	// 配置数据库连接池
	sqlDB, dbError := db.DB()
	if dbError != nil {
		panic(dbError)
	}
	sqlDB.SetConnMaxIdleTime(10 * time.Second)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 将监控数据库对象赋值到 global.MonitorDB 中
	global.MonitorDB = db
	global.LOG.Info("init monitor db successfully")
}
