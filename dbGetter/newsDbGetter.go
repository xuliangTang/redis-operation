package dbGetter

import (
	"fmt"
	"goRedis/db"
	"goRedis/models/newsModel"
	"goRedis/redis"
)

func NewsDbGetter(newsId string) redis.DBGetterFunc {
	return func() interface{} {
		fmt.Println("from db")
		news := newsModel.New()
		db.Db.Where("id=?", newsId).First(&news)
		return news
	}
}
