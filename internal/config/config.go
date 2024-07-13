package config

import (
	"log"

	"github.com/leetcode-golang-classroom/golang-redis-sample/internal/util"
	"github.com/spf13/viper"
)

type Config struct {
	RedisUrl string `mapstructure:"REDIS_URL"`
}

var AppCfg *Config

func init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigType("env")
	v.SetConfigName(".env")
	v.AutomaticEnv()
	// bind environment
	util.FailOnError(v.BindEnv("REDIS_URL"), "failed to bind env REDIS_URL")
	err := v.ReadInConfig()
	if err != nil {
		log.Println("Load from environment variable")
	}
	err = v.Unmarshal(&AppCfg)
	if err != nil {
		util.FailOnError(err, "Failed to read enivronment")
	}
}
