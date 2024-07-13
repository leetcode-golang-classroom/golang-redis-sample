package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/leetcode-golang-classroom/golang-redis-sample/internal/config"
	"github.com/leetcode-golang-classroom/golang-redis-sample/internal/redis"
	"github.com/leetcode-golang-classroom/golang-redis-sample/internal/util"
)

func main() {
	redisURL := config.AppCfg.RedisUrl
	rdb, err := redis.New(redisURL)
	if err != nil {
		util.FailOnError(err, fmt.Sprintf("failed to connect to %s\n", redisURL))
	}
	defer rdb.Close()
	ping, err := rdb.Ping(context.Background())
	if err != nil {
		util.FailOnError(err, fmt.Sprintf("failed to ping to %s\n", redisURL))
	}
	log.Println(ping)
	type Person struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Age        int    `json:"age"`
		Occupation string `json:"occupation"`
	}
	eddieID := uuid.NewString()
	jsonString, err := json.Marshal(Person{
		ID:         eddieID,
		Name:       "eddie",
		Age:        36,
		Occupation: "engineer leader",
	})
	if err != nil {
		util.FailOnError(err, "failed for marshal")
	}
	eddieKey := fmt.Sprintf("eddie:%s", eddieID)
	err = rdb.Set(context.Background(), eddieKey, jsonString, 0)
	if err != nil {
		util.FailOnError(err, fmt.Sprintf("failed to set key %s with value %s\n", eddieKey, jsonString))
	}
	result, err := rdb.Get(context.Background(), eddieKey)
	if err != nil {
		util.FailOnError(err, fmt.Sprintf("failed to get key %s\n", eddieKey))
	}
	log.Println(result)
}
