package services

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/repositories"
	"errors"
	"fmt"
)

type ILocationService interface {
	AllLocations() ([]entities.Location, error)
	LocationByID(idStr string) (*entities.Location, error)
	CreateLocation(location entities.Location) (entities.Location, error)
	DeleteLocation(idStr string) (entities.Location, error)
	UpdateLocation(idStr string, location entities.Location) (entities.Location, error)
}

type LocationService struct {
	Repo            repositories.LocationRepository
	DestinationRepo repositories.DestinationRepository
}

func (service *LocationService) AllLocations() ([]entities.Location, error) {
	locations, err := service.Repo.AllLocations()
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (service *LocationService) LocationByID(idStr string) (*entities.Location, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

	location, err := service.Repo.LocationByID(id)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (service *LocationService) CreateLocation(location entities.Location) (entities.Location, error) {
	location, err := service.Repo.CreateLocation(location)
	if err != nil {
		return entities.Location{}, err
	}
	return location, nil
}

func (service *LocationService) DeleteLocation(idStr string) (entities.Location, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.Location{}, errors.New("invalid ID format")
	}

	if err := service.DestinationRepo.DeleteDestinationsByLocationID(id); err != nil {
		return entities.Location{}, err
	}

	location, err := service.Repo.DeleteLocation(id)
	if err != nil {
		return entities.Location{}, err
	}
	return location, nil
}

func (service *LocationService) UpdateLocation(idStr string, location entities.Location) (entities.Location, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.Location{}, errors.New("invalid ID format")
	}

	location, err := service.Repo.UpdateLocation(id, location)
	if err != nil {
		return entities.Location{}, err
	}
	return location, nil
}
