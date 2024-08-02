package main

import (
	"fmt"
	"net/http"

	"github.com/kauancf/estudo/tree/main/api_students/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", start)
	e.GET("/students", getStudents)
	e.POST("/students", creatStudents)
	e.GET("/students/:id", getStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handlerb
func start(c echo.Context) error {
	return c.String(http.StatusOK, "START")
}

func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "List of all students")
}

func creatStudents(c echo.Context) error {
	db.Addstudent()
	return c.String(http.StatusOK, "Creat Student")
}

func getStudent(c echo.Context) error {
	id := c.Param("id")
	getStudent := fmt.Sprintf("Get %s student", id)
	return c.String(http.StatusOK, getStudent)
}

func updateStudent(c echo.Context) error {
	id := c.Param("id")
	updateStudent := fmt.Sprintf("Update %s student", id)
	return c.String(http.StatusOK, updateStudent)
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStudent := fmt.Sprintf("Delete %s student", id)
	return c.String(http.StatusOK, deleteStudent)
}