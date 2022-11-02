package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Employee struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	TabNum     int64              `json:"tab_num" bson:"tab_num,omitempty"`
	FirstName  string             `json:"first_name" bson:"first_name,omitempty"`
	LastName   string             `json:"last_name" bson:"last_name,omitempty"`
	Age        int64              `json:"age" bson:"age,omitempty"`
	Department *Department        `json:"department" bson:"department,omitempty"`
	Address    *Address           `json:"address" bson:"address,omitempty"`
	Courses    []*Course          `json:"courses" bson:"courses,omitempty"`
	IsActive   bool               `json:"is_active" bson:"is_active,omitempty"`
	IsGpx      bool               `json:"is_gpx" bson:"is_gpx,omitempty"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

type EmployeeBody struct {
	TabNum       int64    `json:"tab_num" binding:"required" `
	FirstName    string   `json:"first_name" binding:"required" `
	LastName     string   `json:"last_name" binding:"required" `
	Age          int64    `json:"age" binding:"required" `
	DepartmentID string   `json:"department_id" binding:"required" `
	Address      *Address `json:"address" binding:"required" `
	Courses      []string `json:"courses" binding:"required" `
	IsActive     bool     `json:"is_active,omitempty"`
	IsGpx        bool     `json:"is_gpx,omitempty"`
}
