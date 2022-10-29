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
		departmentGroup := apiGroup.Group("/departments")
		{
			departmentRepo := repositories.NewDepartmentRepository(mongoDB)
			departmentService := services.NewDepartmentService(departmentRepo)
			departmentController := controllers.NewDepartmentController(departmentService)

			departmentGroup.GET("/all", departmentController.GetDepartmentsList)
			departmentGroup.GET("/:id", departmentController.GetDepartment)
			departmentGroup.POST("/create", departmentController.CreateDepartment)
			departmentGroup.PUT("/update/:id", departmentController.UpdateDepartment)
			departmentGroup.DELETE("/delete/:id", departmentController.DeleteDepartment)
		}

		courseGroup := apiGroup.Group("/courses")
		{
			courseRepo := repositories.NewCourseRepository(mongoDB)
			courseService := services.NewCourseService(courseRepo)
			courseController := controllers.NewCourseController(courseService)

			courseGroup.GET("/all", courseController.GetCoursesList)
			courseGroup.GET("/:id", courseController.GetCourse)
			courseGroup.POST("/create", courseController.CreateCourse)
			courseGroup.PUT("/update/:id", courseController.UpdateCourse)
			courseGroup.DELETE("/delete/:id", courseController.DeleteCourse)
		}

	}
}
