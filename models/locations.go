package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Location struct {
	gorm.Model
	Id   int    `json:"id`
	Name string `json:"name"`
	Text string `json:"text"`
	Deleted_at string `json:"deleted_at"`
}

func GetLocationById(id int) *Location {
	log.Println(id)

	location := &Location{}
	err := GetDB().Table("locations").Where("id = ?", id).First(location).Error
	if err != nil {
		return nil
	}

	return location
}
