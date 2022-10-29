package course

import (
	"github.com/sanzharanarbay/golang_mongoDB_example/internal/models"
	"net/url"
)

type CourseRequest struct {
	model *models.Course
}

func NewCourseRequest(model *models.Course) *CourseRequest {
	return &CourseRequest{
		model: model,
	}
}

func (c *CourseRequest) Validate() url.Values {
	errs := url.Values{}

	if len(c.model.Title) < 3 || len(c.model.Title) > 255 {
		errs.Add("title", "The title field must be between 3-255 chars!")
	}

	if len(c.model.Author) < 3 || len(c.model.Author) > 255 {
		errs.Add("author", "The author field must be between 3-500 chars!")
	}

	if c.model.Hours == 0 {
		errs.Add("hours", "The hours field is required!")
	}

	if c.model.Price == 0 {
		errs.Add("price", "The price field is required!")
	}

	if c.model.Order == 0 {
		errs.Add("order", "The order field is required!")
	}

	return errs
}
