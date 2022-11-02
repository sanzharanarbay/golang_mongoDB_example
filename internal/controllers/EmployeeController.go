package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/models"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/requests/employee"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/services"
	"log"
	"net/http"
)

type EmployeeController struct {
	employeeService *services.EmployeeService
}

func NewEmployeeController(employeeService *services.EmployeeService) *EmployeeController {
	return &EmployeeController{
		employeeService: employeeService,
	}
}

func (e *EmployeeController) CreateEmployee(c *gin.Context) {
	var request models.EmployeeBody

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		log.Println(err)
		return
	}

	validErrs := employee.NewEmployeeRequest(&request).Validate()
	if len(validErrs) > 0 {
		errors := map[string]interface{}{"errors": validErrs}
		c.JSON(http.StatusUnprocessableEntity, errors)
		return
	}

	created, err := e.employeeService.InsertEmployee(&request)
	code := http.StatusCreated

	response := map[string]interface{}{}

	if err == nil {
		log.Println("The employee saved successfully")
		response["status"] = true
		response["message"] = "The employee saved successfully"
		response["insertID"] = created
	} else {
		response["status"] = false
		response["message"] = err.Error()
		code = http.StatusUnprocessableEntity
	}
	c.JSON(code, response)
}

func (e *EmployeeController) GetEmployee(c *gin.Context) {
	var employeeObj *models.Employee
	var err error
	ID := c.Param("id")
	employeeObj, err = e.employeeService.GetEmployeeById(ID)
	if err != nil {
		log.Printf("ERROR - %s", err)
		response := map[string]interface{}{
			"status":  false,
			"message": err.Error(),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	c.JSON(http.StatusOK, employeeObj)
}

func (e *EmployeeController) GetEmployeesList(c *gin.Context) {
	employeeList, err := e.employeeService.GetAllEmployees()
	if err != nil {
		response := map[string]interface{}{
			"status":  false,
			"message": err.Error(),
		}
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	c.JSON(http.StatusOK, employeeList)
}

func (e *EmployeeController) UpdateEmployee(c *gin.Context) {
	var request models.EmployeeBody

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		log.Println(err)
		return
	}

	validErrs := employee.NewEmployeeRequest(&request).Validate()
	if len(validErrs) > 0 {
		errors := map[string]interface{}{"errors": validErrs}
		c.JSON(http.StatusUnprocessableEntity, errors)
		return
	}

	ID := c.Param("id")

	updated, err := e.employeeService.UpdateEmployee(&request, ID)
	code := http.StatusOK

	response := map[string]interface{}{}

	if err == nil {
		log.Println("The employee updated successfully")
		response["status"] = true
		response["data"] = updated
	} else {
		response["status"] = false
		response["message"] = err.Error()
		code = http.StatusUnprocessableEntity
	}

	c.JSON(code, response)
}

func (e *EmployeeController) DeleteEmployee(c *gin.Context) {
	ID := c.Param("id")
	err := e.employeeService.DeleteEmployee(ID)
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
			"message": "The employee deleted successfully",
		}
	}
	c.JSON(code, response)
}
