package middleware

import (
	"LoginService/model"
	"LoginService/serializer"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CurrentUser()  gin.HandlerFunc{


	handlerFunc := func(context *gin.Context) {

		session := sessions.Default(context)

		uid := session.Get("user_id")
		if uid != nil{
			user,err := model.GetUser(uid)
			fmt.Println(user)
			if err == nil{
				context.Set("user", &user)
			}
		}

		context.Next()
	}

	return handlerFunc
}


func AuthRequired() gin.HandlerFunc{
	return func(context *gin.Context) {
		if user , _ := context.Get("user"); user != nil{
			if _,ok := user.(*model.User); ok{
				context.Next()
				return
			}
		}

		context.JSON(200, serializer.Response{
			Status:401,
			Msg:"no login",
		})

		context.Abort()

	}
}