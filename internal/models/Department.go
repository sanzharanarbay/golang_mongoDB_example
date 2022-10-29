package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Department struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" binding:"required" bson:"name,omitempty"`
	ShortName string             `json:"short_name" binding:"required"  bson:"short_name,omitempty"`
	Total     struct {
		Employees int64 `json:"employees" binding:"required" bson:"employees,omitempty"`
		Gpx       int64 `json:"gpx" binding:"required"  bson:"gpx,omitempty"`
	} `json:"total"`
	IsActive  bool      `json:"is_active" binding:"required" bson:"is_active,omitempty"`
	Order     int64     `json:"order" binding:"required" bson:"order,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

