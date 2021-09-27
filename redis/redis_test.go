package redis

import "testing"

func TestNewRedis(t *testing.T) {
	redis := NewRedis(&Config{
		Host: "127.0.0.1",
		Port: "6379",
		Db:   0,
	})
	if redis == nil {
		t.Errorf("redis connect err")
	}
}
