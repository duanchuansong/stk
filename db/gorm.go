package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Host         string
	Port         string
	User         string
	Pass         string
	DbName       string
	MaxIdleConns int
	MaxOpenConns int
}

func NewDb(cfg *Mysql) *gorm.DB {
	dsn := cfg.User + ":" + cfg.Pass + "@tcp(" + cfg.Host + ":"+cfg.Port+")/" + cfg.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("open database err,err:"+err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("open database err,err:"+err.Error())
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxIdleConns)
	return db
}
