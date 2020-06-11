package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type UserBase struct {
	gorm.Model
	Avatar            string    `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`
	IDentityCard      string    `gorm:"column:identity_card;type:varchar(32);not null" json:"identity_card"`
	NikeName          string    `gorm:"column:nike_name;type:varchar(16);not null" json:"nike_name"`
	InviteCode        string    `gorm:"column:invite_code;type:varchar(16);not null" json:"invite_code"`
	CommonLiveCity    string    `gorm:"column:common_live_city;type:varchar(126);not null" json:"common_live_city"`
	Birthday          time.Time `gorm:"column:birthday;type:datetime" json:"birthday"`
	MovmentPrograms   string    `gorm:"column:movment_programs;type:varchar(128);not null" json:"movment_programs"`
	ObjectExpectation string    `gorm:"column:object_expectation;type:varchar(128);not null" json:"object_expectation"`
	SocialAccount     string    `gorm:"column:social_account;type:varchar(32);not null" json:"social_account"`
	WechatAccount     string    `gorm:"column:wechat_account;type:varchar(32);not null" json:"wechat_account"`
	IshideWechat      string    `gorm:"column:ishide_wechat;type:varchar(8);not null" json:"ishide_wechat"`
	Height            float64   `gorm:"column:height;type:decimal(2,1);not null" json:"height"`
	Weight            float64   `gorm:"column:weight;type:decimal(2,0);not null" json:"weight"`
	Introduce         string    `gorm:"column:introduce;type:varchar(128);not null" json:"introduce"`
	EditNum           bool      `gorm:"column:edit_num;type:tinyint(1);not null" json:"edit_num"`
	EditTime          time.Time `gorm:"column:edit_time;type:datetime" json:"edit_time"`
	Carrers           []Carrer  `gorm:"many2many:userbase_carrers;"`
}
type Carrer struct {
	gorm.Model
	CarrerName string    `gorm:"column:carrer_name;type:varchar(16);not null" json:"carrer_name"`
}

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "root:000000@tcp(192.168.204.128:3306)/by?charset=utf8mb4&parseTime=True&loc=Local")
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
	db.DropTableIfExists(&UserBase{})
	db.DropTableIfExists(&Carrer{})
	db.AutoMigrate(&UserBase{}, &Carrer{})

	var l1 = Carrer{}
	l1.CarrerName = "会计"
	db.Create(&l1)

	var l2 = Carrer{}
	l2.CarrerName = "财务"
	db.Create(&l2)

	var l3 = Carrer{}
	l3.CarrerName = "程序员"
	db.Create(&l3)

	u1 := &UserBase{
		Carrers: []Carrer{
			l1,
			l2,
		},
	}
	db.Create(&u1)
	u2 := &UserBase{
		Carrers: []Carrer{
			l2,
			l3,
		},
	}
	db.Create(&u2)
	var userbase UserBase
	db.Find(&userbase, "2")
	var cars []Carrer
	db.Model(&userbase).Related(&cars, "Carrers")
	fmt.Println(userbase)
	fmt.Println(cars)
}
