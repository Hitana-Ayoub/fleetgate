package models

import (
	"github.com/jinzhu/gorm"
)

type Vehicle struct {
	gorm.Model

	Plate    string `json:"plate" gorm:"typevarchar(13); uniqueindex"`
	Brand    string `json:"brand"`
	Driver   string `json:"driver"`
	Status   string `json:"status"`
	Engine   string `json:"engine"`
	KM       int    `json:"km"`
	ImageURL string `json:"imageURL"`
}
