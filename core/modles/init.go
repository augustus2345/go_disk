package modles

import (
	"github.com/go-redis/redis/v8"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine = Init()
var RDB = InitRedis()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/go_disk?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("xorm new Engine failed...,error:%v", err)
		return nil
	}
	return engine
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
