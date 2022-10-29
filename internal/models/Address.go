package models

type Address struct {
	Country   string `json:"country""`
	City      string `json:"city"`
	Region    string `json:"region"`
	Street    string `json:"street"`
	House     string `json:"house,omitempty"`
	Flat      string `json:"flat,omitempty"`
	Floor     string `json:"floor,omitempty"`
	PostIndex string `json:"post_index,omitempty"`
}
