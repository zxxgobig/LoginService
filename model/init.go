package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

const maxIdleConn  = 20 //空闲连接数
const maxOpenConn = 100 //打开的连接数
const maxLifeTime = time.Second * 30 //超时时间

func Database(conn string)  {
	db, err := gorm.Open("mysql", conn)

	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(maxIdleConn)
	db.DB().SetMaxOpenConns(maxOpenConn)
	db.DB().SetConnMaxLifetime(maxLifeTime)

	DB = db

}
