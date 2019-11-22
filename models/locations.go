package models

import (
	"github.com/jinzhu/gorm"
)

type Location struct {
	gorm.Model
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Text       string  `json:"text"`
	Deleted_at string  `json:"deleted_at"`
	Image_url  string  `json:"image_url"`
	Lat        float64 `json:"lat"`
	Long       float64 `json:"long"`
}

func GetAllLocations() []*Location {

	locations := make([]*Location, 0)
	err := GetDB().Table("locations").Find(&locations).Error
	if err != nil {
		return nil
	}

	return locations
}
func GetLocationById(id int) *Location {

	location := &Location{}
	err := GetDB().Table("locations").Where("id = ?", id).First(location).Error
	if err != nil {
		return nil
	}

	return location
}
