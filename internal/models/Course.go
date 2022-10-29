package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Course struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" binding:"required" bson:"title,omitempty"`
	Hours     int64              `json:"hours" binding:"required" bson:"hours,omitempty"`
	Author    string             `json:"author" binding:"required" bson:"author,omitempty"`
	Price     float64            `json:"price" binding:"required" bson:"price,omitempty"`
	Order     int64              `json:"order" binding:"required" bson:"order,omitempty"`
	IsActive  bool               `json:"is_active" binding:"required" bson:"is_active,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
