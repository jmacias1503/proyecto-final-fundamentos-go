package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Id   int `gorm:"primaryKey" autoIncrement:"true"`
	Name string
}

// Estructura de estudiantes
type Student struct {
	Id    int    `json:"student_id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Group string `json:"group"`
	Email string `json:"email"`
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to DB")
	}
	db.AutoMigrate(&Subject{}, &Student{})
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
			"total": count,
			"list":  subjects,
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

	//---- Student functions ----
	//Get students
	router.GET("api/students", func(ctx *gin.Context) {
		var count int64
		var students []Student
		db.Find(&students)
		db.Model(&students).Count(&count)
		ctx.HTML(200, "count-template.html", gin.H{
			"title": "Estudiantes registrados",
			"total": count,
			"list":  students,
		})
	})
	//Create student
	router.POST("api/students", func(ctx *gin.Context) {
		var student Student
		if ctx.BindJSON(&student) == nil {
			db.Create(&student)
			ctx.JSON(200, student)
		} else {
			ctx.JSON(400, gin.H{
				"error": "Invalid payload",
			})
		}
	})
	students := []Student{}
	//Delete student
	router.DELETE("api/students/:student_id", func(ctx *gin.Context) {
		id := ctx.Param("student_id")
		for i, student := range students {
			if strconv.Itoa(student.Id) == id { // Estará bien esto???
				students = append(students[:i], students[i+1])
				db.Delete(&students)
				ctx.JSON(200, gin.H{
					"message": "User deleted.",
				})
				return
			}
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User not found.",
		})
	})
	//Get singular student
	router.GET("api/students/:student_id", func(ctx *gin.Context) {
		id := ctx.Param("student_id")
		for i, student := range students {
			if strconv.Itoa(student.Id) == id {
				students[i] = student
				db.Find(&student)
				ctx.JSON(200, student)
			}
		}
	})
	// Update student
	router.PUT("api/students/:student_id", func(ctx *gin.Context) {
		id := ctx.Param("studen_id")
		type Body struct {
			Name  string `json:"name"`
			Group string `json:"group"`
			Email string `json:"email"`
		}
		var body Body
		err := ctx.BindJSON(&body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid payload",
			})
		} else {
			for i, student := range students {
				if strconv.Itoa(student.Id) == id {
					students[i].Name = body.Name
					students[i].Group = body.Group
					students[i].Email = body.Email
					// Como se actualiza en base de datos?
					ctx.JSON(http.StatusAccepted, students[i])
					return
				}
			}
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
		}
	})
	router.Run(":8000")
}
