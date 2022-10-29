package department

import (
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/models"
	"net/url"
)

type DepartmentRequest struct {
	model *models.Department
}

func NewDepartmentRequest(model *models.Department) *DepartmentRequest {
	return &DepartmentRequest{
		model: model,
	}
}

func (d *DepartmentRequest) Validate() url.Values {
	errs := url.Values{}

	if len(d.model.Name) < 3 || len(d.model.Name) > 255 {
		errs.Add("name", "The name field must be between 3-255 chars!")
	}

	if len(d.model.ShortName) < 3 || len(d.model.ShortName) > 255 {
		errs.Add("short_name", "The short_name field must be between 3-500 chars!")
	}

	if d.model.Total.Gpx == 0 {
		errs.Add("total.gpx", "The total.gpx field is required!")
	}

	if d.model.Total.Employees == 0 {
		errs.Add("total.employees", "The total.employees field is required!")
	}

	if d.model.Order == 0 {
		errs.Add("order", "The order field is required!")
	}

	return errs
}
