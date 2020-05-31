package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// User 包含多个 emails, UserID 为外键
type User struct {
	gorm.Model
	Emails   []Email
}

type Email struct {
	gorm.Model
	Email   string
	UserID  uint
}

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:000000@tcp(192.168.163.133:3306)/ex4?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error(), "9999")
		panic("连接数据库失败")
		return
	}
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Duration(300) * time.Second)
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	db.SingularTable(true)
	db.LogMode(true)
	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&Email{})
	db.AutoMigrate(&User{}, &Email{})
	var userx = User{}
	db.Create(&userx)
	var usery = User{}
	db.Create(&usery)

	var email1 = Email{}
	email1.UserID = 7
	email1.Email = "emailxxx"
	db.Create(&email1)
	var email2 = Email{}
	email2.UserID = 7
	email2.Email = "emailyyy"
	db.Create(&email2)

	var user User
	user.ID = 7
	var email []Email
	db.Model(&user).Related(&email)
	fmt.Println(user.ID, user.Emails)
	fmt.Println(email)
}
