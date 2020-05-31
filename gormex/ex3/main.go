package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// User 包含一个 CreditCard, UserID 为外键
type User struct {
	gorm.Model
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:000000@tcp(192.168.163.133:3306)/ex3?charset=utf8mb4&parseTime=True&loc=Local")
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
	db.DropTableIfExists(&CreditCard{})
	db.AutoMigrate(&User{}, &CreditCard{})
	var userx = User{}
	db.Create(&userx)
	var usery = User{}
	db.Create(&usery)

	var card1 = CreditCard{}
	card1.UserID = 1
	card1.Number = "111"
	db.Create(&card1)
	var card2 = CreditCard{}
	card2.UserID = 2
	card2.Number = "222"
	db.Create(&card2)

	var usertmp User
	usertmp.ID = 2
	var card CreditCard
	db.Model(&usertmp).Related(&card)
	fmt.Println(usertmp.ID, usertmp.CreditCard)
	fmt.Println(card.ID, card.Number, card.UserID)
}
