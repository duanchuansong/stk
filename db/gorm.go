package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DbName string
}

func NewDb(cfg *Mysql) *gorm.DB {
	dsn := cfg.User + ":" + cfg.Pass + "@tcp(" + cfg.Host + ":3306)/" + cfg.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	return db
}
