package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Tb_UserBase struct {
	gorm.Model
	RegisterTime      time.Time     `gorm:"column:register_time;type:datetime" json:"register_time"`
	Avatar            string        `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`
	IDentityCard      string        `gorm:"column:identity_card;type:varchar(32);not null" json:"identity_card"`
	NikeName          string        `gorm:"column:nike_name;type:varchar(16);not null" json:"nike_name"`
	InviteCode        string        `gorm:"column:invite_code;type:varchar(16);not null" json:"invite_code"`
	CommonLiveCity    string        `gorm:"column:common_live_city;type:varchar(126);not null" json:"common_live_city"`
	Birthday          time.Time     `gorm:"column:birthday;type:datetime" json:"birthday"`
	MovmentPrograms   string        `gorm:"column:movment_programs;type:varchar(128);not null" json:"movment_programs"`
	ObjectExpectation string        `gorm:"column:object_expectation;type:varchar(128);not null" json:"object_expectation"`
	SocialAccount     string        `gorm:"column:social_account;type:varchar(32);not null" json:"social_account"`
	WechatAccount     string        `gorm:"column:wechat_account;type:varchar(32);not null" json:"wechat_account"`
	IshideWechat      string        `gorm:"column:ishide_wechat;type:varchar(8);not null" json:"ishide_wechat"`
	Height            float64       `gorm:"column:height;type:decimal(2,1);not null" json:"height"`
	Weight            float64       `gorm:"column:weight;type:decimal(2,0);not null" json:"weight"`
	Introduce         string        `gorm:"column:introduce;type:varchar(128);not null" json:"introduce"`
	EditNum           bool          `gorm:"column:edit_num;type:tinyint(1);not null" json:"edit_num"`
	EditTime          time.Time     `gorm:"column:edit_time;type:datetime" json:"edit_time"`
	Carrers           []Tb_Carrer   `gorm:"many2many:userbase_carrers;"`
	Favorites         []Tb_Favorite `gorm:"many2many:userbase_favorites;"`
	Blacks            []Tb_Black    `gorm:"many2many:userbase_blacks;"`

	PhotoAlbum []Tb_PhotoAlbum
	Secret     string  `gorm:"column:secret;type:varchar(16);not null" json:"secret"`
	Money      float64 `gorm:"column:money;type:decimal(16,1);not null" json:"money"`
}
type Tb_PhotoAlbum struct {
	gorm.Model
	PhotoName     string
	Tb_UserBaseID uint
}
type Tb_Carrer struct {
	gorm.Model
	CarrerName string `gorm:"column:carrer_name;type:varchar(16);not null" json:"carrer_name"`
}
type Tb_Favorite struct {
	gorm.Model
	FavoriteName string `gorm:"column:favorite_name;type:varchar(16);not null" json:"favorite_name"`
}
type Tb_Black struct {
	gorm.Model
	BlackName string `gorm:"column:black_name;type:varchar(16);not null" json:"black_name"`
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
	db.DropTableIfExists(&Tb_UserBase{})
	db.DropTableIfExists(&Tb_Carrer{})
	db.DropTableIfExists(&Tb_Favorite{})
	db.DropTableIfExists(&Tb_Black{})
	db.DropTableIfExists(&Tb_PhotoAlbum{})
	db.AutoMigrate(&Tb_UserBase{}, &Tb_Carrer{}, &Tb_Favorite{}, &Tb_Black{}, &Tb_PhotoAlbum{})
	//------------------------------------------------------------_
	var l1 = Tb_Carrer{}
	l1.CarrerName = "ca111"
	db.Create(&l1)

	var l2 = Tb_Carrer{}
	l2.CarrerName = "ca222"
	db.Create(&l2)

	var l3 = Tb_Carrer{}
	l3.CarrerName = "ca333"
	db.Create(&l3)
	//------------------------------------------------------------_
	var f1 = Tb_Favorite{}
	f1.FavoriteName = "fff111"
	db.Create(&f1)

	var f2 = Tb_Favorite{}
	f2.FavoriteName = "fff222"
	db.Create(&f2)

	var f3 = Tb_Favorite{}
	f3.FavoriteName = "fff333"
	db.Create(&f3)
	//------------------------------------------------------------_
	var b1 = Tb_Black{}
	b1.BlackName = "b11111"
	db.Create(&b1)

	var b2 = Tb_Black{}
	b2.BlackName = "b22222"
	db.Create(&b2)

	var b3 = Tb_Black{}
	b3.BlackName = "b33333"
	db.Create(&b3)
	//------------------------------------------------------------_
	var ph1 = Tb_PhotoAlbum{}
	ph1.PhotoName = "ph111"
	ph1.Tb_UserBaseID = 1
	db.Create(&ph1)

	var ph2 = Tb_PhotoAlbum{}
	ph2.PhotoName = "ph222"
	ph2.Tb_UserBaseID = 2
	db.Create(&ph2)

	var ph3 = Tb_PhotoAlbum{}
	ph3.PhotoName = "ph333"
	ph3.Tb_UserBaseID = 2
	db.Create(&ph3)

	u1 := &Tb_UserBase{
		Carrers: []Tb_Carrer{
			l2,
			l3,
		},
		Favorites: []Tb_Favorite{
			f2,
			f3,
		},
		Blacks: []Tb_Black{
			b2,
			b3,
		},
	}
	db.Create(&u1)
	u2 := &Tb_UserBase{
		Carrers: []Tb_Carrer{
			l1,
			l3,
		},
		Favorites: []Tb_Favorite{
			f1,
			f3,
		},
		Blacks: []Tb_Black{
			b1,
			b3,
		},
	}
	db.Create(&u2)
	var userbase Tb_UserBase
	db.Find(&userbase, 2)
	var cars []Tb_Carrer
	var favs []Tb_Favorite
	var blacks []Tb_Black
	var phs []Tb_PhotoAlbum
	db.Model(&userbase).Related(&cars, "Carrers").Related(&favs, "Favorites").Related(&blacks, "Blacks").Related(&phs)
	fmt.Println("Carrers", cars, " Favorites", favs, "***", blacks, "$$$:", phs)
}
