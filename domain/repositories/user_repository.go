package repositories

import "Trip-Trove-API/domain/entities"

type UserRepository interface {
	AllUsers() ([]entities.User, error)
	AllUserIDs() ([]uint, error)
	UserByID(id uint) (*entities.User, error)
	Register(user entities.User) (entities.User, error)
	Login(loginData entities.LoginRequest) (entities.LoginResponse, error)
	UpdateUser(id uint, updatedUser entities.User) (entities.User, error)
	DeleteUser(id uint) (entities.User, error)
}
