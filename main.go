package main

import (
	"github.com/gin-gonic/gin"
	"goRedis/cachePool"
	"goRedis/dbGetter"
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

	r := gin.Default()
	r.GET("/news/:id", func(context *gin.Context) {
		newsId := context.Param("id")
		newsCache := cachePool.NewsCache()
		defer cachePool.NewsCacheRelease(newsCache)
		newsCache.DBGetter = dbGetter.NewsDbGetter(newsId)
		context.Header("Content-Type", "application/json")
		context.String(200, newsCache.GetCache("news"+newsId).(string))
	})
	r.Run(":8080")



}