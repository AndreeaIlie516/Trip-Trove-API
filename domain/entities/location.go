package entities

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name        string `gorm:"column:name;not null;unique" json:"name" validate:"required,min=3,max=30"`
	Country     string `gorm:"column:country;not null" json:"country" validate:"required"`
	Description string `gorm:"column:description" json:"description" validate:"max=256"`
}
