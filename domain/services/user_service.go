package services

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/repositories"
	"errors"
	"fmt"
)

type UserService struct {
	Repo repositories.UserRepository
}

func (service *UserService) AllUsers() ([]entities.User, error) {
	users, err := service.Repo.AllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (service *UserService) UserByID(idStr string) (*entities.User, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

	user, err := service.Repo.UserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserService) Register(user entities.User) (entities.User, error) {
	user, err := service.Repo.Register(user)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (service *UserService) Login(loginData entities.LoginRequest) (entities.LoginResponse, error) {
	response, err := service.Repo.Login(loginData)
	if err != nil {
		return entities.LoginResponse{}, err
	}
	return response, nil
}

func (service *UserService) DeleteUser(idStr string) (entities.User, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.User{}, errors.New("invalid ID format")
	}

	user, err := service.Repo.DeleteUser(id)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (service *UserService) UpdateUser(idStr string, user entities.User) (entities.User, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.User{}, errors.New("invalid ID format")
	}

	user, err := service.Repo.UpdateUser(id, user)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}
