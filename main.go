package main

import (
	"encoding/json"
	"fmt"
	"goRedis/db"
	"goRedis/models/newsModel"
	"goRedis/redis"
	"log"
	"time"
)

func main()  {
	//getName := redis.NewStringOperation().Set("name","txl",
	//	redis.WithExpire(time.Second*30),
	//	redis.WithNX(),
	//	redis.WithXX(),
	//).Unwrap()
	//fmt.Println(getName)

	//getValues := redis.NewStringOperation().MGet("name", "age", "sex").Iterator()
	//for getValues.HasNext() {
	//	fmt.Println(getValues.Next())
	//}

	newsId := 3
	newsCache := redis.NewStringCache(redis.NewStringOperation(), time.Second*30)
	newsCache.DBGetter = func() string {
		fmt.Println("from db")
		news := newsModel.New()
		db.Db.Where("id=?", newsId).First(&news)
		bs,err := json.Marshal(news)
		if err != nil {
			log.Fatal(err)
		}
		return string(bs)
	}
	fmt.Println(newsCache.GetCache("test"))

}