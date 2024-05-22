package repositories

import "Trip-Trove-API/domain/entities"

type DestinationRepository interface {
	AllDestinations() ([]entities.Destination, error)
	AllDestinationIDs() ([]uint, error)
	DestinationByID(id uint) (*entities.Destination, error)
	DestinationIDsForLocation(locationID uint) ([]uint, error)
	DeleteDestinationsByLocationID(locationID uint) error
	CreateDestination(destination entities.Destination) (entities.Destination, error)
	UpdateDestination(id uint, updatedDestination entities.Destination) (entities.Destination, error)
	DeleteDestination(id uint) (entities.Destination, error)
}
