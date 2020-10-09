package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const Active string = "active"
const (
	PasswordCost int = 12
)

type User struct {
	gorm.Model
	UserName string
	PasswordDigest string
	NickName string
	Status string
	Avatar string `gorm:"size:1000"`

}

func (user *User) TableName() string{
	return "user"
}

func GetUser(id interface{})  (User, error){
	var user User
	result := DB.Model(user).First(&user, id)

	return user, result.Error
}

func (user *User) CreateUser()  error {

	return DB.Model(user).Create(user).Error

}
func (user *User) CheckUserName(name string) error {

	return DB.Model(user).Where("user_name = ?", name).First(&user).Error
}

func (user *User) CheckNickName(name string)  int{
	var count int
	DB.Model(user).Where("nick_name = ?", name).Count(&count)

	return count
}


func (user *User) CheckPassword(password string) bool{
	fmt.Println(password, user.PasswordDigest)
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}