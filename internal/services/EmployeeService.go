package services

import (
	"errors"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/models"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/repositories"
)

type EmployeeService struct {
	employeeRepository   *repositories.EmployeeRepository
	courseRepository     *repositories.CourseRepository
	departmentRepository *repositories.DepartmentRepository
}

func NewEmployeeService(employeeRepository *repositories.EmployeeRepository, courseRepository *repositories.CourseRepository,
	departmentRepository *repositories.DepartmentRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepository:   employeeRepository,
		courseRepository:     courseRepository,
		departmentRepository: departmentRepository,
	}
}

func (e *EmployeeService) InsertEmployee(employeeBody *models.EmployeeBody) (interface{}, error) {
	var courses []*models.Course
	var courseObj *models.Course
	var departmentObj *models.Department
	departmentObj, _ = e.departmentRepository.GetDepartment(employeeBody.DepartmentID)
	if departmentObj == nil {
		return nil, errors.New("no department with that Id exists")
	}

	for _, element := range employeeBody.Courses {
		courseObj, _ = e.courseRepository.GetCourse(element)
		if courseObj == nil {
			return nil, errors.New("no course with that Id exists")
		}
		courses = append(courses, courseObj)
	}

	check, _ := e.employeeRepository.CheckEmployeeExistence(employeeBody.TabNum)
	if check != nil{
		return nil, errors.New("the Employee already exists")
	}

	var employeeObj models.Employee
	employeeObj.TabNum = employeeBody.TabNum
	employeeObj.FirstName = employeeBody.FirstName
	employeeObj.LastName = employeeBody.LastName
	employeeObj.Age = employeeBody.Age
	employeeObj.Department = departmentObj
	employeeObj.Address = employeeBody.Address
	employeeObj.Courses = courses
	employeeObj.IsActive = employeeBody.IsActive
	employeeObj.IsGpx = employeeBody.IsGpx

	state, err := e.employeeRepository.SaveEmployee(&employeeObj)
	return state, err
}

func (e *EmployeeService) GetEmployeeById(ID string) (*models.Employee, error) {
	state, err := e.employeeRepository.GetEmployee(ID)
	return state, err
}

func (e *EmployeeService) GetAllEmployees() (*[]models.Employee, error) {
	state, err := e.employeeRepository.GetAllEmployees()
	return state, err
}

func (e *EmployeeService) UpdateEmployee(employeeBody *models.EmployeeBody, ID string) (interface{}, error) {
	var courses []*models.Course
	var courseObj *models.Course
	var departmentObj *models.Department
	departmentObj, _ = e.departmentRepository.GetDepartment(employeeBody.DepartmentID)
	if departmentObj == nil {
		return nil, errors.New("no department with that Id exists")
	}

	for _, element := range employeeBody.Courses {
		courseObj, _ = e.courseRepository.GetCourse(element)
		if courseObj == nil {
			return nil, errors.New("no course with that Id exists")
		}
		courses = append(courses, courseObj)
	}


	var employeeObj models.Employee
	employeeObj.TabNum = employeeBody.TabNum
	employeeObj.FirstName = employeeBody.FirstName
	employeeObj.LastName = employeeBody.LastName
	employeeObj.Age = employeeBody.Age
	employeeObj.Department = departmentObj
	employeeObj.Address = employeeBody.Address
	employeeObj.Courses = courses
	employeeObj.IsActive = employeeBody.IsActive
	employeeObj.IsGpx = employeeBody.IsGpx

	state, err := e.employeeRepository.UpdateEmployee(&employeeObj, ID)
	return state, err
}

func (e *EmployeeService) DeleteEmployee(ID string) error {
	err := e.employeeRepository.DeleteEmployee(ID)
	return err
}
