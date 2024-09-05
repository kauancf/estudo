package api

import (
	"errors"
	"net/http"

	"strconv"

	"github.com/kauancf/estudo/tree/main/api_students/schemas"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Handlerb

// getStudents godoc
//
//	@Summary        Get a list of students
//	@Description    Retrieve students details
//	@Tags           students
//	@Accept         json
//	@Produce        json
//	@Param          register path    int  false    "Registration"
//	@Success        200 {object} schemas.StudentResponse
//	@Failure        404
//	@Router         /students [get]
func start(c echo.Context) error {
	return c.String(http.StatusOK, "START")
}

func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}

	active := c.QueryParam("active")

	if active != "" {
		act, err := strconv.ParseBool(active)
		if err != nil {
			log.Error().Err(err).Msgf("[api] error to parse boolean")
			return c.String(http.StatusInternalServerError, "Failed to parse boolean")
		}
		students, err = api.DB.GetFilteredStudents(act)
		if err != nil {
			log.Error().Err(err).Msgf("[api] error to Filtered active")
		}
	}

	listOfStudents := map[string][]schemas.StudentResponse{"students:": schemas.NewResponse(students)}

	return c.JSON(http.StatusOK, listOfStudents)
}

// createStudent godoc
//
//	@Summary        Create student
//	@Description    Create student
//	@Tags           students
//	@Accept         json
//	@Produce        json
//	@Success        200 {object} schemas.StudentResponse
//	@Failure        400
//	@Router         /students [post]
func (api *API) creatStudents(c echo.Context) error {
	studentReq := StudentRequest{}
	if err := c.Bind(&studentReq); err != nil {
		return err
	}

	if err := studentReq.Validate(); err != nil {
		log.Error().Err(err).Msgf("[api] error validating struct")
		return c.String(http.StatusBadRequest, "Error validating  Student")

	}

	student := schemas.Student{
		Name:   studentReq.Name,
		CPF:    studentReq.CPF,
		Email:  studentReq.Email,
		Age:    studentReq.Age,
		Active: *studentReq.Active,
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to creat Student")
	}

	return c.JSON(http.StatusOK, student)
}

// getStudent godoc
//
//	@Summary        Get student by ID
//	@Description    Retrieve student details using ID
//	@Tags           students
//	@Accept         json
//	@Produce        json
//	@Success        200 {object} schemas.StudentResponse
//	@Failure        404
//	@Failure        500
//	@Router         /students/{id} [get]
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

// updateStudent godoc
//
//	@Summary        Update Student
//	@Description    Update student details
//	@Tags           students
//	@Accept         json
//	@Produce        json
//	@Success        200 {object} schemas.StudentResponse
//	@Failure        404
//	@Failure        500
//	@Router         /students/{id} [put]
func (api *API) updateStudent(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student ID")
	}

	receivedStudent := schemas.Student{}
	if err := c.Bind(&receivedStudent); err != nil {
		return err
	}

	updateStudent, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	student := updateStudentInfo(receivedStudent, updateStudent)

	if err := api.DB.UpdateStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save student")
	}

	return c.JSON(http.StatusOK, student)
}

// deleteStudent godoc
//
//	@Summary        Delete Student
//	@Description    Delete student details
//	@Tags           students
//	@Accept         json
//	@Produce        json
//	@Success        200 {object} schemas.StudentResponse
//	@Failure        404
//	@Failure        500
//	@Router         /students/{id} [delete]
func (api *API) deleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student ID")
	}

	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	if err := api.DB.DeletStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delet student")
	}

	return c.JSON(http.StatusOK, student)
}

func updateStudentInfo(receivedStudent, student schemas.Student) schemas.Student {
	if receivedStudent.Name != "" {
		student.Name = receivedStudent.Name
	}
	if receivedStudent.CPF > 0 {
		student.CPF = receivedStudent.CPF
	}
	if receivedStudent.Email != "" {
		student.Email = receivedStudent.Email
	}
	if receivedStudent.Age > 0 {
		student.Age = receivedStudent.Age
	}
	if receivedStudent.Active != student.Active {
		student.Active = receivedStudent.Active
	}
	return student
}
