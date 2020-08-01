package model

import (
	"fmt"
	"log"
	"time"

	"go_web/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	args := fmt.Sprintf("%s:%s@/go_web?charset=utf8&parseTime=True&loc=Local", config.DbName, config.DbPassword)
	db, err := gorm.Open("mysql", args)
	if err != nil {
		log.Fatal("数据库连接错误", err)
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	// 赋值成全局变量
	DB = db

	fmt.Println("连接成功")
	// 自动迁移数据库
	DB.AutoMigrate(&User{}, &Role{}, &Permission{})
}
