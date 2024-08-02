package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	CPF    int
	Email  string
	Age    int
	Active bool
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
}

func Addstudent() {
	db := Init()
	student := Student{
		Name:   "Kauan",
		CPF:    987654321,
		Email:  "Kauanf@gmail.com",
		Age:    30,
		Active: true,
	}

	// result := db.Create(&student)  //Forma completa
	// if result.Error != nil {
	// 	fmt.Println("Erro to create student")
	// }
	// fmt.Println("Create student")

	if result := db.Create(&student); result.Error != nil { // Forma simplificada
		fmt.Println("Erro to create student")
	}
	fmt.Println("Create student")

}
