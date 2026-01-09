package persistence

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"admin-pro/internal/config"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	
	// 设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	
	log.Println("Database connection established")
	return db, nil
}
