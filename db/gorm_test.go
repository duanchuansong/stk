package db

import "testing"

func TestNewDb(t *testing.T) {
	db := NewDb(&Mysql{
		Host:         "127.0.0.1",
		Port:         "3306",
		User:         "root",
		Pass:         "admin123",
		DbName:       "tms",
		MaxIdleConns: 10,
		MaxOpenConns: 200,
	})
	if db.Error != nil {
		t.Errorf("new db failed,err:%v",db.Error)
	}
}
