package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// `User`属于`Human`, `HumanID`为外键

type User struct {
	gorm.Model
	Human   Human
	HumanID int
}

type Human struct {
	gorm.Model
	Number string
}

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:000000@tcp(192.168.163.133:3306)/ex2?charset=utf8mb4&parseTime=True&loc=Local")
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
	db.DropTableIfExists(&Human{})
	db.AutoMigrate(&User{}, &Human{})
	var userx = User{HumanID:1}
	var usery = User{HumanID:2}
	var humanx = Human{Number: "human1"}
	var humany = Human{Number: "human2"}

	db.Create(&userx)
	db.Create(&usery)
	db.Create(&humanx)
	db.Create(&humany)

	var user User
	user.HumanID = 2
	var hum Human
	db.Model(&user).Related(&hum)
	fmt.Println(user)
	fmt.Println(hum)

	//// SELECT * FROM credit_cards WHERE user_id = 123; // 123 is user's primary key
	// CreditCard是user的字段名称，这意味着获得user的CreditCard关系并将其填充到变量
	// 如果字段名与变量的类型名相同，如上例所示，可以省略，如：
	//db.Model(&user).Related(&card)

}
