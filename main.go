package main

import (
	"LoginService/conf"
	"LoginService/server"
)

func main()  {
	conf.Init()
	//fmt.Println(model.DB, cache.RedisClient)

	r := server.NewRouter()
	r.Run(":2000")
}
