package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// User 包含并属于多个 languages, 使用 `user_languages` 表连接
type User struct {
	gorm.Model
	Languages         []Language `gorm:"many2many:ds_user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}


var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:000000@tcp(192.168.163.133:3306)/ex5?charset=utf8mb4&parseTime=True&loc=Local")
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
	db.DropTableIfExists(&Language{})
	db.AutoMigrate(&User{}, &Language{})

	var l1 = Language{}
	l1.Name = "Chinese"
	db.Create(&l1)
	var l2 = Language{}
	l2.Name = "English"
	db.Create(&l2)
	var l3 = Language{}
	l3.Name = "Japanese"
	db.Create(&l3)

	u1 := &User{
		Languages: []Language{
			l1,
			l2,
		},
	}
	db.Create(&u1)
	u2 := &User{
		Languages: []Language{
			l2,
			l3,
		},
	}
	db.Create(&u2)

	var user User
	db.Find(&user, 2)
	var languages []Language
	db.Model(&user).Related(&languages, "Languages")
	fmt.Println(user)
	fmt.Println(languages)
}
