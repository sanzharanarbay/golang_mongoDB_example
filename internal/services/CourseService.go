package services

import (
	"errors"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/models"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/repositories"
)

type CourseService struct {
	courseRepository *repositories.CourseRepository
}

func NewCourseService(courseRepository *repositories.CourseRepository) *CourseService {
	return &CourseService{
		courseRepository: courseRepository,
	}
}

func (c *CourseService) InsertCourse(course *models.Course) (interface{}, error) {
	check, _ := c.courseRepository.CheckCourseExistence(course.Title, course.Author)
	if check != nil{
		return nil, errors.New("the Course already exists")
	}
	state, err := c.courseRepository.SaveCourse(course)
	return state, err
}

func (c *CourseService) GetCourseById(ID string) (*models.Course, error) {
	state, err := c.courseRepository.GetCourse(ID)
	return state, err
}

func (c *CourseService) GetAllCourses() (*[]models.Course, error) {
	state, err := c.courseRepository.GetAllCourses()
	return state, err
}

func (c *CourseService) UpdateCourse(course *models.Course, ID string) (interface{}, error) {
	state, err := c.courseRepository.UpdateCourse(course, ID)
	return state, err
}

func (c *CourseService) DeleteCourse(ID string) error {
	err := c.courseRepository.DeleteCourse(ID)
	return err
}
