package api

import (
	"fmt"
	"net/http"

	"github.com/kauancf/estudo/tree/main/api_students/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

func NewServer() *API {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentDB := db.NewStudentHandler(database)

	return &API{
		Echo: e,
		DB:   studentDB,
	}

}

func (api *API) ConfigureRoutes() {
	// Routes
	api.Echo.GET("/", start)
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.creatStudents)
	api.Echo.GET("/students/:id", api.getStudent)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
}

func (api *API) Start() error {
	// Start server
	return api.Echo.Start(":8080")
}

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
	id := c.Param("id")
	getStudent := fmt.Sprintf("Get %s student", id)
	return c.String(http.StatusOK, getStudent)
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
