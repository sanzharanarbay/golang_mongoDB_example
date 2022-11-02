package models

type Address struct {
	Country   string `json:"country" binding:"required" bson:"country,omitempty"`
	City      string `json:"city" binding:"required" bson:"city,omitempty"`
	Region    string `json:"region" binding:"required" bson:"region,omitempty"`
	Street    string `json:"street" binding:"required" bson:"street,omitempty"`
	House     string `json:"house,omitempty" bson:"house,omitempty"`
	Flat      string `json:"flat,omitempty" bson:"flat,omitempty"`
	Floor     string `json:"floor,omitempty" bson:"floor,omitempty"`
	PostIndex string `json:"post_index,omitempty" bson:"post_index,omitempty"`
}
