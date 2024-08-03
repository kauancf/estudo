package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"registration"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
}

func AddStudent(student Student) error {
	db := Init()

	if result := db.Create(&student); result.Error != nil { // Forma simplificada
		return result.Error
	}
	fmt.Println("Create student")
	return nil

	// result := db.Create(&student)  //Forma completa
	// if result.Error != nil {
	// 	fmt.Println("Erro to create student")
	// }
	// fmt.Println("Create student")
}

func GetStudents() ([]Student, error) {
	students := []Student{}

	db := Init()
	err := db.Find(&students).Error
	return students, err
}
