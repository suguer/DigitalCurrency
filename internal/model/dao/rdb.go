package dao

import (
	"DigitalCurrency/internal/config"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var Rdb *redis.Client

func RedisInit(conf *config.RedisConf) {
	options := redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			conf.Host,
			conf.Port), // Redis地址
		DB:          conf.Database,                   // Redis库
		PoolSize:    5,                               // Redis连接池大小
		MaxRetries:  3,                               // 最大重试次数
		IdleTimeout: time.Second * time.Duration(10), // 空闲链接超时时间
	}
	if viper.GetString("redis_passwd") != "" {
		options.Password = viper.GetString("redis_passwd")
	}
	Rdb = redis.NewClient(&options)
	/*
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
			pong, err := Rdb.Ping(ctx).Result()
			if err == redis.Nil {
				log.Sugar.Debug("[store_redis] Nil reply returned by Rdb when key does not exist.")
			} else if err != nil {
				color.Red.Printf("[store_redis] redis connRdb err,err=%s", err)
				panic(err)
			} else {
				log.Sugar.Debug("[store_redis] redis connRdb success,suc=%s", pong)
			}*/
}
