package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() {
	// viper config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/machhub/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	// exit if config file not found
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// set env prefix
	viper.SetEnvPrefix("MACHHUB")

	// set default values
	viper.SetDefault("app.workdir", "web")
	viper.SetDefault("laravel.appkeysize", 32)
	viper.SetDefault("app.port", 8080)
	viper.SetDefault("db.type", "mysql")
	viper.SetDefault("db.prefix", "web_")
	viper.SetDefault("db.port", "3306")
	viper.SetDefault("db.host", "127.0.0.1")
	viper.SetDefault("aws.mailport", "2587")
	viper.SetDefault("aws.mailencryption", "tls")
	viper.SetDefault("db.admindb", "admin")
	viper.SetDefault("app.company", "Machhub ApS")
	viper.SetDefault("app.datadir", "data")
	viper.SetDefault("app.timezone", "Europe/Copenhagen")
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)

	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	// check if app.apitoken is empty
	if viper.GetString("app.apitoken") == "" {
		// exit if app.apitoken is empty
		log.Fatal("app.apitoken can't be empty")
	}
}

func Get(key string) interface{} {
	return viper.Get(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
}
