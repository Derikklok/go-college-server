package student

import (
	"net/http"
	"strconv"

	"github.com/Derikklok/go-college-server/internal/shared"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

// fn name		take service as an input return - a controller
func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

// tells create student function belongs to this controller
// c -> similar to this keyword
func (c *Controller) CreateStudent(ctx *gin.Context) {

	var req CreateStudentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {

		ctx.JSON(http.StatusBadRequest, shared.ErrorResponse{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})

		return
	}

	student, err := c.service.CreateStudent(req)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, shared.ErrorResponse{
			Success: false,
			Message: "Failed to create student",
			Error:   err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, shared.SuccessResponse{
		Success: true,
		Data:    student,
	})
}

func (c *Controller) GetAllStudents(ctx *gin.Context) {

	students, err := c.service.GetAllStudents()

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, shared.ErrorResponse{
			Success: false,
			Message: "Failed to fetch students",
			Error:   err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, shared.SuccessResponse{
		Success: true,
		Data:    students,
	})
}

func (c *Controller) DeleteStudent(ctx *gin.Context) {

	idParam := ctx.Param("id")

	id, _ := strconv.Atoi(idParam)

	err := c.service.DeleteStudent(uint(id))

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, shared.ErrorResponse{
			Success: false,
			Message: "Delete failed",
			Error:   err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, shared.SuccessResponse{
		Success: true,
		Data:    "Student deleted",
	})
}
