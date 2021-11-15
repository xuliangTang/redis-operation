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
		db := db.Db.Where("id=?", newsId).Find(&news)
		if db.Error!=nil || db.RowsAffected == 0 {
			return nil
		}
		return news
	}
}
