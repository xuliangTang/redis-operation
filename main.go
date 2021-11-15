package main

import (
	"github.com/gin-gonic/gin"
	"goRedis/cachePool"
	"goRedis/common"
	"goRedis/dbGetter"
	"goRedis/models/newsModel"
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
	r.Use(common.ErrorHandler())
	r.GET("/news/:id", func(context *gin.Context) {
		// 获取参数
		newsId := context.Param("id")

		// 从对象池获取新闻缓存对象
		newsCache := cachePool.NewsCache()
		defer cachePool.NewsCacheRelease(newsCache)

		// 设置DBGetter
		newsCache.DBGetter = dbGetter.NewsDbGetter(newsId)

		// 取缓存输出，缓存中没有会调用上面的DBGetter
		news := newsModel.New()
		// context.Header("Content-Type", "application/json")
		newsCache.GetCacheForObject("news"+newsId, news)
		context.JSON(200, news)
	})
	r.Run(":8080")



}