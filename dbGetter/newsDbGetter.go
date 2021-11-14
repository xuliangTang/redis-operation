package dbGetter

import (
	"encoding/json"
	"fmt"
	"goRedis/db"
	"goRedis/models/newsModel"
	"goRedis/redis"
	"log"
)

func NewsDbGetter(newsId string) redis.DBGetterFunc {
	return func() string {
		fmt.Println("from db")
		news := newsModel.New()
		db.Db.Where("id=?", newsId).First(&news)
		bs,err := json.Marshal(news)
		if err != nil {
			log.Fatal(err)
		}
		return string(bs)
	}
}
