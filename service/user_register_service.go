package service

import (
	"LoginService/model"
	"LoginService/serializer"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserRegisterService struct {
	NickName string `form:"nickname" json:"nickname" binding:"required,min=5,max=30"`
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=6,max=40"`
}

func (service *UserRegisterService)Valid()  *serializer.Response {
	//两次密码不一致
	if service.Password != service.PasswordConfirm{
		return &serializer.Response{
			Status: 402,
			Msg:"确认密码不一致",
		}
	}
	var user model.User

	//昵称被占用
	if count := user.CheckNickName(service.NickName);count >0{
		return &serializer.Response{
			Status:402,
			Msg:"昵称已存在",
		}
	}



	//用户名已注册
	if err := user.CheckUserName(service.UserName); err != nil {
		panic(err)
		fmt.Println(err)
		return &serializer.Response{
			Status:402,
			Msg:"用户名已注册",
		}
	}

	return nil
}

func (service *UserRegisterService) Register()  (model.User, *serializer.Response){
	user := model.User{
		NickName:service.NickName,
		UserName:service.UserName,
		Status : model.Active,
	}

	err := service.Valid()
	if err != nil{
		return user,err
	}

	//密码加密
	bytes,err2 := bcrypt.GenerateFromPassword([]byte(service.Password),model.PasswordCost)
	if err2 != nil{
		return user, &serializer.Response{
			Status: 402,
			Msg:"密码加密失败",
		}
	}
	user.PasswordDigest = string(bytes)

	if err3 := user.CreateUser(); err3 != nil{
		return user, &serializer.Response{
			Status:402,
			Msg:"注册失败",
		}
	}

	return user,nil
}