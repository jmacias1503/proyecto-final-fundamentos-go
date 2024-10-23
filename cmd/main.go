package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Id int `gorm:primaryKey;autoIncrement:true`
	Name string
}


func main() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to DB")
	}
	db.AutoMigrate(&Subject{})
	router := gin.Default()
	fmt.Println("running app")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{
			"title": "Sistema de Control escolar de la GUAK",
		})
	})

	router.GET("/subjects", func(ctx *gin.Context) {
		var count int64
		var subjects []Subject
		db.Find(&subjects)
		db.Model(&subjects).Count(&count)
		ctx.HTML(200, "count-template.html", gin.H{
			"title": "Materias registradas",
			"total" : count,
			"list": subjects,
		})
	})

	router.POST("/api/subjects", func(ctx *gin.Context) {
		var subject Subject
		if ctx.BindJSON(&subject) == nil {
			db.Create(&subject)
			ctx.JSON(201, subject)
		} else {
			ctx.JSON(400, gin.H{
				"error": "Invalid payload",
			})
		}
	})

	router.Run(":8000")
}
