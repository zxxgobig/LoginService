package server

import (
	"LoginService/api"
	"LoginService/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func NewRouter() *gin.Engine{
	r := gin.Default()

	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))//中间件
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	v1 := r.Group("api/v1")
	{
		fmt.Println("first？")
	}
	{
		v1.GET("ping", api.Ping)

		//登录
		v1.POST("user/login",api.UserLogin)
		//注册
		v1.POST("user/register", api.UserRegister)


		v1.Use(middleware.AuthRequired())
		{
			fmt.Println("auth first ？")
		}
		{
			v1.GET("user/me", api.UserMe)

		}

	}
	{
		fmt.Println("next？")
	}




	return r
}
