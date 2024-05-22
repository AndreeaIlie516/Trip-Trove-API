package entities

import "gorm.io/gorm"

type Destination struct {
	gorm.Model
	Name             string `gorm:"column:name;not null;unique" json:"name" validate:"required,min=3,max=50"`
	LocationID       uint   `gorm:"column:location_id;not null" json:"location_id" validate:"required,number"`
	ImageUrl         string `gorm:"column:image_url" json:"image_url" validate:"max=100"`
	Description      string `gorm:"column:description" json:"description" validate:"min=10,max=256"`
	VisitorsLastYear int    `gorm:"column:visitors_last_year" json:"visitors_last_year" validate:"gte=0"`
	IsPrivate        bool   `gorm:"column:is_private;not null" json:"is_private"`
}
