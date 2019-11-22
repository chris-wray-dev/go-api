package db

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type location struct {
	gorm.Model
	id         int     `gorm:"UNIQUE PRIMARY KEY"`
	name       string  `gorm:"type:VARCHAR(30) UNIQUE NOT NULL"`
	text       string  `gorm:"type:TEXT NOT NULL"`
	image_url  string  `gorm:"type:varchar(50) NOT NULL"`
	lat        float64 `gorm:"type:FLOAT NOT NULL"`
	long       float64 `gorm:"type:FLOAT NOT NULL"`
	deleted_at string  `gorm:"type:TEXT"`
}

func CreateLocationsObject() []location {
	locations := []location{
		{
			id:         6,
			name:       `Hacienda`,
			text:       `Some text here`,
			image_url:  `URL TBC`,
			lat:        2,
			long:       3,
			deleted_at: ``,
		},
	}

	return locations

}
