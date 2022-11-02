package v1

import (
	"github.com/gin-gonic/gin"
	mongo_db "github.com/sanzharanarbay/golang_mongoDB_example/internal/configs/mongo-db"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/controllers"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/repositories"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/services"
)

func ApiRoutes(prefix string, router *gin.Engine) {
	mongoDB := mongo_db.NewMongoDB()
	apiGroup := router.Group(prefix)
	{

		// departments
		departmentRepo := repositories.NewDepartmentRepository(mongoDB)
		departmentService := services.NewDepartmentService(departmentRepo)
		departmentController := controllers.NewDepartmentController(departmentService)
		departmentGroup := apiGroup.Group("/departments")
		{
			departmentGroup.GET("/all", departmentController.GetDepartmentsList)
			departmentGroup.GET("/:id", departmentController.GetDepartment)
			departmentGroup.POST("/create", departmentController.CreateDepartment)
			departmentGroup.PUT("/update/:id", departmentController.UpdateDepartment)
			departmentGroup.DELETE("/delete/:id", departmentController.DeleteDepartment)
		}

		// courses
		courseRepo := repositories.NewCourseRepository(mongoDB)
		courseService := services.NewCourseService(courseRepo)
		courseController := controllers.NewCourseController(courseService)
		courseGroup := apiGroup.Group("/courses")
		{
			courseGroup.GET("/all", courseController.GetCoursesList)
			courseGroup.GET("/:id", courseController.GetCourse)
			courseGroup.POST("/create", courseController.CreateCourse)
			courseGroup.PUT("/update/:id", courseController.UpdateCourse)
			courseGroup.DELETE("/delete/:id", courseController.DeleteCourse)
		}

		// employees
		employeeRepo := repositories.NewEmployeeRepository(mongoDB)
		employeeService := services.NewEmployeeService(employeeRepo, courseRepo, departmentRepo)
		employeeController := controllers.NewEmployeeController(employeeService)
		employeeGroup := apiGroup.Group("/employees")
		{
			employeeGroup.GET("/all", employeeController.GetEmployeesList)
			employeeGroup.GET("/:id", employeeController.GetEmployee)
			employeeGroup.POST("/create", employeeController.CreateEmployee)
			employeeGroup.PUT("/update/:id", employeeController.UpdateEmployee)
			employeeGroup.DELETE("/delete/:id", employeeController.DeleteEmployee)
		}


	}
}
