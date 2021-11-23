package main

import (
	"ex1/recycle/global"
	_ "ex1/recycle/init"
	"ex1/recycle/model"
	"github.com/gin-gonic/gin"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		result := GetAllCategories()
		c.JSON(200, gin.H{
			"message": result,
		})
	})
	r.Run()
}

func GetAllCategories() []model.Category {
	var categorys []model.Category
	global.Gdb.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	return categorys
}
