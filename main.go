package main

import (
	"github.com/karolpiernikarz/automanage/cache"
	"github.com/karolpiernikarz/automanage/config"
	"github.com/karolpiernikarz/automanage/helpers"
	newAdmin "github.com/karolpiernikarz/automanage/services/newadmin"
	"github.com/karolpiernikarz/automanage/services/redis"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	config.InitLogger()
	helpers.Init()
	cache.Init()
	config.InitCron()
	redis.Init(viper.GetString("redis.host"), viper.GetString("redis.port"), viper.GetString("redis.password"), viper.GetInt("redis.db"))
	newAdmin.InitDB()
	config.InitRoutes()
}
