package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var rdb *redis.Client
var ctx = context.Background()

func Init() (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"), // 密码
		DB:       viper.GetInt("redis.db"),          // 数据库
		PoolSize: viper.GetInt("redis.conn_pool"),   // 连接池大小
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
func Close() {
	_ = rdb.Close()
}
