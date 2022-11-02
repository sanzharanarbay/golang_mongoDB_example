package employee

import (
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/models"
	"net/url"
)

type EmployeeRequest struct {
	model *models.EmployeeBody
}

func NewEmployeeRequest(model *models.EmployeeBody) *EmployeeRequest {
	return &EmployeeRequest{
		model: model,
	}
}

func (e *EmployeeRequest) Validate() url.Values {
	errs := url.Values{}

	if len(e.model.FirstName) < 3 || len(e.model.FirstName) > 255 {
		errs.Add("first_name", "The first_name field must be between 3-255 chars!")
	}

	if len(e.model.LastName) < 3 || len(e.model.LastName) > 255 {
		errs.Add("last_name", "The last_name field must be between 3-500 chars!")
	}

	if e.model.TabNum == 0 {
		errs.Add("tab_num", "The tab_num field is required!")
	}

	if e.model.Age == 0 {
		errs.Add("age", "The age field is required!")
	}

	if len(e.model.DepartmentID) < 3 || len(e.model.DepartmentID) > 30 {
		errs.Add("department_id", "The department_id field must be between 3-30 chars!")
	}

	if len(e.model.Address.Country) < 3 || len(e.model.Address.Country) > 255 {
		errs.Add("address.country", "The address.country field must be between 3-255 chars!")
	}

	if len(e.model.Address.City) < 3 || len(e.model.Address.City) > 255 {
		errs.Add("address.city", "The address.city field must be between 3-500 chars!")
	}

	if len(e.model.Address.Region) < 3 || len(e.model.Address.Region) > 255 {
		errs.Add("address.region", "The address.region field must be between 3-255 chars!")
	}

	if len(e.model.Address.Street) < 3 || len(e.model.Address.Street) > 255 {
		errs.Add("address.street", "The address.street field must be between 3-500 chars!")
	}

	return errs
}
