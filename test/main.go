package main

import (
	"HiChat/global"
	"context"
	"fmt"
	"time"
)

func main() {
	/*dsn := "root:li244088167@tcp(localhost:3306)/hichat?charset=utf8mb4&parseTime=True&loc=Local"

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		err = db.AutoMigrate(&models.Message{})
		if err != nil {
			panic(err)
		}
	}*/
	/*opt := redis.Options{
		Addr:     "180.163.78.203:6379", // redis地址
		Password: "",                    // redis密码，没有则留空
		DB:       0,                     // 默认数据库，默认是0
	}
	db := redis.NewClient(&opt)*/
	result, err := global.RedisDB.Set(context.Background(), "name", time.Second, time.Second*10).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
