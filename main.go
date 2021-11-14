package main

import (
	"fmt"
	"goRedis/redis"
	"time"
)

func main()  {
	getName := redis.NewStringOperation().Set("name","txl",
		redis.WithExpire(time.Second*30),
		redis.WithNx(),
	).Unwrap()
	fmt.Println(getName)

	//getValues := redis.NewStringOperation().MGet("name", "age", "sex").Iterator()
	//for getValues.HasNext() {
	//	fmt.Println(getValues.Next())
	//}

}