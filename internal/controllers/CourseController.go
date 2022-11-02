package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/models"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/requests/course"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/services"
	"log"
	"net/http"
)

type CourseController struct {
	courseService *services.CourseService
}

func NewCourseController(courseService *services.CourseService) *CourseController {
	return &CourseController{
		courseService: courseService,
	}
}

func (c *CourseController) CreateCourse(ctx *gin.Context) {
	var request models.Course

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "invalid json")
		log.Println(err)
		return
	}

	validErrs := course.NewCourseRequest(&request).Validate()
	if len(validErrs) > 0 {
		errors := map[string]interface{}{"errors": validErrs}
		ctx.JSON(http.StatusUnprocessableEntity, errors)
		return
	}

	created, err := c.courseService.InsertCourse(&request)
	code := http.StatusCreated

	response := map[string]interface{}{}

	if err == nil {
		fmt.Println("The course saved successfully")
		response["status"] = true
		response["message"] = "The course saved successfully"
		response["insertID"] = created
	}else{
		response["status"] = false
		response["message"] = err.Error()
		code = http.StatusUnprocessableEntity
	}
	ctx.JSON(code, response)
}

func (c *CourseController) GetCourse(ctx *gin.Context) {
	var courseObj *models.Course
	var err error
	ID := ctx.Param("id")
	courseObj, err = c.courseService.GetCourseById(ID)
	if err != nil {
		log.Printf("ERROR - %s", err)
		response := map[string]interface{}{
			"status":  false,
			"message": err.Error(),
		}
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	ctx.JSON(http.StatusOK, courseObj)
}

func (c *CourseController) GetCoursesList(ctx *gin.Context) {
	coursesList, err := c.courseService.GetAllCourses()
	if err != nil {
		response := map[string]interface{}{
			"status":  false,
			"message": err.Error(),
		}
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	ctx.JSON(http.StatusOK, coursesList)
}

func (c *CourseController) UpdateCourse(ctx *gin.Context) {
	var request models.Course

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "invalid json")
		log.Println(err)
		return
	}

	validErrs := course.NewCourseRequest(&request).Validate()
	if len(validErrs) > 0 {
		errors := map[string]interface{}{"errors": validErrs}
		ctx.JSON(http.StatusUnprocessableEntity, errors)
		return
	}

	ID := ctx.Param("id")

	updated, err := c.courseService.UpdateCourse(&request, ID)
	code := http.StatusOK

	response := map[string]interface{}{}

	if err == nil {
		fmt.Println("The course updated successfully")
		response["status"] = true
		response["data"] = updated
	} else {
		response["status"] = false
		response["message"] = err.Error()
		code = http.StatusUnprocessableEntity
	}

	ctx.JSON(code, response)
}

func (c *CourseController) DeleteCourse(ctx *gin.Context) {
	ID := ctx.Param("id")
	err := c.courseService.DeleteCourse(ID)
	code := http.StatusOK
	response := map[string]interface{}{}
	if err != nil {
		log.Printf("ERROR - %s", err)
		response = map[string]interface{}{
			"status":  false,
			"message": err.Error(),
		}
		code = http.StatusUnprocessableEntity
	} else {
		response = map[string]interface{}{
			"status":  true,
			"message": "The course deleted successfully",
		}
		fmt.Println("The course deleted successfully")
	}
	ctx.JSON(code, response)
}
