package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/models"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/requests/department"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/services"
	"log"
	"net/http"
)

type DepartmentController struct {
	departmentService *services.DepartmentService
}

func NewDepartmentController(departmentService *services.DepartmentService) *DepartmentController {
	return &DepartmentController{
		departmentService: departmentService,
	}
}

func (d *DepartmentController) CreateDepartment(c *gin.Context) {
	var request models.Department

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		log.Println(err)
		return
	}

	validErrs := department.NewDepartmentRequest(&request).Validate()
	if len(validErrs) > 0 {
		errors := map[string]interface{}{"errors": validErrs}
		c.JSON(http.StatusUnprocessableEntity, errors)
		return
	}

	created, err := d.departmentService.InsertDepartment(&request)
	code := http.StatusCreated
	response := map[string]interface{}{}

	if err == nil {
		fmt.Println("The department saved successfully")
		response["status"] = true
		response["message"] = "The department saved successfully"
		response["insertID"] = created
	}else{
		response["status"] = false
		response["message"] = err.Error()
	}
	c.JSON(code, response)
}

func (d *DepartmentController) GetDepartment(c *gin.Context) {
	var post *models.Department
	var err error
	ID := c.Param("id")
	post, err = d.departmentService.GetDepartmentById(ID)
	if err != nil {
		log.Printf("ERROR - %s", err)
		response := map[string]interface{}{
			"status":  false,
			"message": err.Error(),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	c.JSON(http.StatusOK, post)
}

func (d *DepartmentController) GetDepartmentsList(c *gin.Context) {
	departmentsList, err := d.departmentService.GetAllDepartments()
	if err != nil {
		response := map[string]interface{}{
			"status":  false,
			"message": err.Error(),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	c.JSON(http.StatusOK, departmentsList)
}

func (d *DepartmentController) UpdateDepartment(c *gin.Context) {
	var request models.Department

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		log.Println(err)
		return
	}

	validErrs := department.NewDepartmentRequest(&request).Validate()
	if len(validErrs) > 0 {
		errors := map[string]interface{}{"errors": validErrs}
		c.JSON(http.StatusUnprocessableEntity, errors)
		return
	}

	ID := c.Param("id")

	updated, err := d.departmentService.UpdateDepartment(&request, ID)
	code := http.StatusOK

	response := map[string]interface{}{}

	if err == nil {
		fmt.Println("The department updated successfully")
		response["status"] = true
		response["data"] = updated
	} else {
		response["status"] = false
		response["message"] = err.Error()
		code = http.StatusUnprocessableEntity
	}

	c.JSON(code, response)
}

func (d *DepartmentController) DeleteDepartment(c *gin.Context) {
	ID := c.Param("id")
	err := d.departmentService.DeleteDepartment(ID)
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
			"message": "The department deleted successfully",
		}
		fmt.Println("The department deleted successfully")
	}
	c.JSON(code, response)
}
