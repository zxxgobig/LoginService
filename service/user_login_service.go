package service

import (
	"LoginService/model"
	"LoginService/serializer"
)

type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=40"`
}

func (service *UserLoginService) Login() (model.User, *serializer.Response) {
	var user model.User
	err := user.CheckUserName(service.UserName)
	if err !=nil{
		return user, &serializer.Response{
			Status:401,
			Msg:"账号不存在",
		}
	}

	if user.CheckPassword(service.Password) == false{
		return user ,&serializer.Response{
			Status:401,
			Msg:"账号密码错误",
		}
	}

	return user,nil
}