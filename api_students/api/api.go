package api

import (
	"github.com/kauancf/estudo/tree/main/api_students/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/kauancf/estudo/tree/main/api_students/docs"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

// @title Student API
// @version 1.0
// @description This is a sample server Student API
// @host localhost:8080
// @BasePath /
// @schemes http
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
	api.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (api *API) Start() error {
	// Start server
	return api.Echo.Start(":8080")
}
