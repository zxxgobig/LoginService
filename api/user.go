package api

import (
	"LoginService/model"
	"LoginService/serializer"
	service2 "LoginService/service"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

func UserMe(ctx *gin.Context){
	user := CurrentUser(ctx)
	res := serializer.BuildUserResponse(*user)
	ctx.JSON(200, res)
}
func UserMe2(ctx *gin.Context){
	//user := CurrentUser(ctx)

	user := model.User{
		UserName:"zxx",
		PasswordDigest:"xxxxxxx",
		NickName:"zheng",
		Status:"1",
		Avatar:"http://ww",
	}
	var data []interface{}
	data = append(data, &user)
	s := serializer.Response{
		Status:0,
		Data: data,
		Msg:"",
		Error:"",
	}
	ctx.JSON(200, s)
}


func UserLogout(ctx *gin.Context){
	s := sessions.Default(ctx)
	s.Clear()
	s.Save()
	ctx.JSON(200, serializer.Response{
		Status:0,
		Msg:"登出成功",
	})
}

func UserLogin(ctx *gin.Context)  {
	var service service2.UserLoginService

	if err := ctx.ShouldBind(&service);err == nil{
		if user,err := service.Login(); err!= nil{
			ctx.JSON(200, err)
		}else {
			s := sessions.Default(ctx)
			s.Clear()
			s.Set("user_id", "")
			s.Save()

			res := serializer.BuildUserResponse(user)
			ctx.JSON(200, res)
		}

	}else {
		ctx.JSON(200,ErrorResponse(err))
	}
	//ctx.JSON(200,ErrorResponse(errors.New("login error")))
}

func UserRegister(ctx *gin.Context){
	var service service2.UserRegisterService
	if err := ctx.ShouldBind(&service);err == nil{
		if user,err := service.Register(); err != nil{
			ctx.JSON(200, err)
		}else {
			res := serializer.BuildUserResponse(user)
			ctx.JSON(200, res)
		}

	}else {
		ctx.JSON(200, ErrorResponse(err))
	}
}


func ErrorResponse(err error) serializer.Response{

	if value,ok := err.(validator.ValidationErrors);ok{
		for _,e := range value{
			field := fmt.Sprintf("Field.%s", e.Field)
			tag := fmt.Sprintf("Tag.Valid.%s", e.Tag)

			return serializer.Response{
				Status:402,
				Msg:fmt.Sprintf("%s%s", field, tag),
				Error:fmt.Sprint(err),

			}
		}
	}

	if _,ok := err.(*json.UnmarshalTypeError);ok{
		return serializer.Response{
			Status:402,
			Msg:"json类型不匹配",
			Error:fmt.Sprint(err),
		}
	}


	return serializer.Response{
		Status:402,
		Msg:"参数错误",
		Error: fmt.Sprint(err),
	}

}