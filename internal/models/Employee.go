package models

type Employee struct {
	TabNum     string     `json:"tab_num"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Age        int64      `json:"age"`
	Department Department `json:"department"`
	Address    Address    `json:"address"`
	Courses    []Course   `json:"courses"`
	IsActive   bool       `json:"is_active"`
	IsGpx      bool       `json:"is_gpx"`
	CreatedAt  string     `json:"created_at"`
	UpdatedAt  string     `json:"updated_at"`
}
