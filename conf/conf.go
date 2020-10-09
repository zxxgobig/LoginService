package conf

import (
	"LoginService/cache"
	"LoginService/model"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Init()  {

	env := os.Getenv("APP_ENV")


	envName := "conf/env/.env."

	switch env {
		case "dev":
			envName += "dev"
		case "test":
			envName += "test"
		case "pro":
			envName += "pro"
		default:
			envName += "dev"
	}
	fmt.Println(envName)

	godotenv.Load(envName)

	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil{
		panic(err)
	}

	conn := os.Getenv("MYSQL_DSN")

	if conn =="" {
		panic("err mysql connect ")
	}
	fmt.Println(conn)

	model.Database(conn)//return DB

	cache.Redis()//return RedisClient


}