package redis

import (
	"github.com/go-redis/redis"
)

type Config struct {
	Host string
	Port string
	Db   int
}

func NewRedis(cfg *Config) (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Host+":"+cfg.Port,
		Password: "", // no password set
		DB:       cfg.Db,  // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		panic("redis connect error,err:"+err.Error())
	}
	return
}
