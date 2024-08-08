package api

import (
	"errors"
	"fmt"
	"net/http"

	"strconv"

	"github.com/kauancf/estudo/tree/main/api_students/db"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// Handlerb

func start(c echo.Context) error {
	return c.String(http.StatusOK, "START")
}

func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}
	return c.JSON(http.StatusOK, students)
}

func (api *API) creatStudents(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to creat Student")
	}

	return c.String(http.StatusOK, "Creat Student")
}

func (api *API) getStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student ID")
	}

	student, err := api.DB.GetStudent(id)
	//não encontrar um studant com esse id (Status not found, 404)
	// algum problema para encontrar um student

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	return c.JSON(http.StatusOK, student)
}

func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	updateStudent := fmt.Sprintf("Update %s student", id)
	return c.String(http.StatusOK, updateStudent)
}

func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStudent := fmt.Sprintf("Delete %s student", id)
	return c.String(http.StatusOK, deleteStudent)
}
