package infrastructure

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

func InitRedis(path string) (*redis.Client, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	addr := fmt.Sprintf("%s:%s", viper.GetString("host"), viper.GetString("port"))
	pass := viper.GetString("password")
	db := viper.GetInt("db")

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
