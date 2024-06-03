package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string     `gorm:"column:username;unique;not null" json:"username" binding:"required" validate:"required,usernameValidator"`
	Password    string     `gorm:"column:password;not null" json:"password" binding:"required" validate:"required"`
	Email       string     `gorm:"column:email;unique;not null" json:"email" binding:"required" validate:"required,email"`
	FirstName   string     `gorm:"column:first_name;not null" json:"first_name" binding:"required" validate:"required,nameValidator"`
	LastName    string     `gorm:"column:last_name;not null" json:"last_name" binding:"required" validate:"required,nameValidator"`
	PhoneNumber string     `gorm:"column:phone_number;unique;not null" json:"phone_number" binding:"required" validate:"required,e164"`
	DateOfBirth string     `gorm:"column:date_of_birth" json:"date_of_birth" binding:"required" validate:"max=20"`
	Address     string     `gorm:"column:address" json:"address" binding:"required" validate:"max=100"`
	Role        AccessType `gorm:"column:access_type,type:tinyint;not null"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Jwt   string `json:"jwt"`
	ID    uint   `json:"id"`
}

type AccessType uint8

const (
	NormalUser AccessType = iota
	Manager
	Admin
)
