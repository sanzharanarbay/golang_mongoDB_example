package services

import (
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/models"
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/repositories"
)

type DepartmentService struct {
	departmentRepository *repositories.DepartmentRepository
}

func NewDepartmentService(departmentRepository *repositories.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		departmentRepository: departmentRepository,
	}
}

func (d *DepartmentService) InsertDepartment(department *models.Department) (interface{}, error) {
	state, err := d.departmentRepository.SaveDepartment(department)
	return state, err
}

func (d *DepartmentService) GetDepartmentById(ID string) (*models.Department, error) {
	state, err := d.departmentRepository.GetDepartment(ID)
	return state, err
}

func (d *DepartmentService) GetAllDepartments() (*[]models.Department, error) {
	state, err := d.departmentRepository.GetAllDepartments()
	return state, err
}

func (d *DepartmentService) UpdateDepartment(department *models.Department, ID string) (interface{}, error) {
	state, err := d.departmentRepository.UpdateDepartment(department, ID)
	return state, err
}

func (d *DepartmentService) DeleteDepartment(ID string) error {
	err := d.departmentRepository.DeleteDepartment(ID)
	return err
}
