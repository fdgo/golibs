package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gormex/ex1/model"
	"log"
	"time"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open("mysql", "root:000000@tcp(192.168.163.133:3306)/ex1?charset=utf8mb4&parseTime=True&loc=Local")
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

	if len(tables) > 0 {
		for _, m := range tables {
			if !db.HasTable(m) {
				err := db.CreateTable(m).Error
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		db.AutoMigrate(tables...)
	}
}
var tables = []interface{}{
	&model.GormexUser{},
	&model.GormexEmail{},
	&model.GormexAddress{},
	&model.GormexLanguage{},
	&model.GormexCreditCard{},
}

func Insert(user model.GormexUser) error {
	return db.Create(&model.GormexUser{Age:user.Age,Birthday: user.Birthday}).Error
}
func QueryOne() (user model.GormexUser,err error) {
	err = db.Where("age = ?", "66").Find(&user).Error
	return user,err
}
func QueryMany()(users []*model.GormexUser,err error)  {
	err = db.Where("age = ?", "66").Find(&users).Error
	return users,err
}
func Update() error {
	return db.Model(&model.GormexUser{}).Where("age=?",13).Update("age", 66).Error
}
func Delete() error {
	return db.Where("age=?",66).Delete(&model.GormexUser{}).Error
}
func main() {
	err := Insert(model.GormexUser{Age: 12,Name:"fred",Birthday: time.Now()})
	fmt.Println(err.Error(),"kkk")
	err2 := Insert(model.GormexUser{Age: 13,Name:"fred",Birthday: time.Now()})
	fmt.Println(err2.Error(),"aaa")

	//
	//db.Create(&model.GormexUser{Age:12,Birthday: time.Now()})
	//db.Create(&model.GormexUser{Age:13,Birthday: time.Now()})
	//db.Create(&model.GormexUser{Age:14,Birthday: time.Now()})
	//
	//db.Model(&model.GormexUser{}).Where("age=?",13).Update("age", 66)
	//db.Where("age=?",66).Delete(&model.GormexUser{})
	//
	//var user []model.GormexUser
	//db.Where("age = ?", "66").Find(&user)
	//fmt.Println(len(user))
}
