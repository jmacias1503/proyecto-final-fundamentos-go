package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Grade struct {
	gorm.Model
	StudentID uint    `json:"student_id" gorm:"not null"`
	SubjectID uint    `json:"subject_id" gorm:"not null"`
	Grade     string  `json:"grade"`
	Student   Student `gorm:"foreignKey:StudentID"`
	Subject   Subject `gorm:"foreignKey:SubjectID"`
}

type Subject struct {
	gorm.Model
	Name  string  `json:"name"`
	Grades []Grade `gorm:"foreignKey:SubjectID"`
}

type Student struct {
	gorm.Model
	Name  string  `json:"name"`
	Group string  `json:"group"`
	Email string  `json:"email"`
	Grades []Grade `gorm:"foreignKey:StudentID"`
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error connecting to DB")
	}
	db.AutoMigrate(&Subject{}, &Student{}, &Grade{})

	router := gin.Default()
	fmt.Println("Running app")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Sistema de Control escolar de la GUAK",
		})
	})

	router.GET("/api/subjects", func(ctx *gin.Context) {
		var subjects []Subject
		db.Find(&subjects)
		ctx.JSON(http.StatusOK, subjects)
	})

	router.POST("/api/subjects", func(ctx *gin.Context) {
		var subject Subject
		if err := ctx.ShouldBindJSON(&subject); err == nil {
			db.Create(&subject)
			ctx.JSON(http.StatusCreated, subject)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		}
	})

	router.GET("/api/subjects/:id", func(ctx *gin.Context) {
		var subject Subject
		id := ctx.Param("id")
		if err := db.First(&subject, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Subject not found."})
		} else {
			ctx.JSON(http.StatusOK, subject)
		}
	})

	router.PUT("/api/subjects/:id", func(ctx *gin.Context) {
		var subject Subject
		id := ctx.Param("id")
		if err := db.First(&subject, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Subject not found."})
			return
		}
		if err := ctx.ShouldBindJSON(&subject); err == nil {
			db.Save(&subject)
			ctx.JSON(http.StatusOK, subject)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		}
	})

	router.DELETE("/api/subjects/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if err := db.Delete(&Subject{}, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Subject not found."})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Subject deleted."})
		}
	})

	router.GET("/api/students", func(ctx *gin.Context) {
		var students []Student
		db.Find(&students)
		ctx.JSON(http.StatusOK, students)
	})

	router.POST("/api/students", func(ctx *gin.Context) {
		var student Student
		if err := ctx.ShouldBindJSON(&student); err == nil {
			db.Create(&student)
			ctx.JSON(http.StatusCreated, student)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		}
	})

	router.GET("/api/students/:id", func(ctx *gin.Context) {
		var student Student
		id := ctx.Param("id")
		if err := db.First(&student, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Student not found."})
		} else {
			ctx.JSON(http.StatusOK, student)
		}
	})

	router.PUT("/api/students/:id", func(ctx *gin.Context) {
		var student Student
		id := ctx.Param("id")
		if err := db.First(&student, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Student not found."})
			return
		}
		if err := ctx.ShouldBindJSON(&student); err == nil {
			db.Save(&student)
			ctx.JSON(http.StatusOK, student)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		}
	})

	router.DELETE("/api/students/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if err := db.Delete(&Student{}, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Student not found."})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Student deleted."})
		}
	})

	router.GET("/api/grades", func(ctx *gin.Context) {
		var grades []Grade
		db.Preload("Student").Preload("Subject").Find(&grades)
		ctx.JSON(http.StatusOK, grades)
	})

	router.POST("/api/grades", func(ctx *gin.Context) {
		var grade Grade
		if err := ctx.ShouldBindJSON(&grade); err == nil {
			db.Create(&grade)
			ctx.JSON(http.StatusCreated, grade)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		}
	})

	router.GET("/api/grades/:id", func(ctx *gin.Context) {
		var grade Grade
		id := ctx.Param("id")
		if err := db.First(&grade, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Grade not found."})
		} else {
			ctx.JSON(http.StatusOK, grade)
		}
	})

	router.PUT("/api/grades/:id", func(ctx *gin.Context) {
		var grade Grade
		id := ctx.Param("id")
		if err := db.First(&grade, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Grade not found."})
			return
		}
		if err := ctx.ShouldBindJSON(&grade); err == nil {
			db.Save(&grade)
			ctx.JSON(http.StatusOK, grade)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		}
	})

	router.DELETE("/api/grades/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if err := db.Delete(&Grade{}, id).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Grade not found."})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Grade deleted."})
		}
	})

	router.Run(":8000")
}
