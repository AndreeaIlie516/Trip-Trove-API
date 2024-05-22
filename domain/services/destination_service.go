package services

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/repositories"
	"Trip-Trove-API/infrastructure/websocket"
	"errors"
	"fmt"
	"github.com/jaswdr/faker"
	"time"
)

type IDestinationService interface {
	AllDestinations() ([]entities.Destination, error)
	DestinationByID(idStr string) (*entities.Destination, error)
	DestinationsByLocationID(locationIDStr string) (*DestinationsByLocation, error)
	CreateDestination(destination entities.Destination) (entities.Destination, error)
	UpdateDestination(idStr string, updatedDestination entities.Destination) (entities.Destination, error)
	DeleteDestination(idStr string) (entities.Destination, error)
	StartGeneratingDestinations(interval time.Duration, f faker.Faker)
	StopGeneratingDestinations()
	GenerateFakeDestination(f faker.Faker) (entities.Destination, error)
}

type DestinationService struct {
	Repo         repositories.DestinationRepository
	LocationRepo repositories.LocationRepository
	Ticker       *time.Ticker
	StopChan     chan bool
	WsManager    *websocket.WebSocketManager
}

type DestinationDetails struct {
	Destination entities.Destination
	Location    entities.Location
}

type DestinationWithLocation struct {
	Destination entities.Destination
	Location    entities.Location
}

type DestinationsByLocation struct {
	Location     entities.Location
	Destinations []entities.Destination
}

var _ IDestinationService = &DestinationService{}

func (service *DestinationService) AllDestinations() ([]entities.Destination, error) {
	destinations, err := service.Repo.AllDestinations()
	if err != nil {
		return nil, err
	}
	return destinations, nil
}

func (service *DestinationService) DestinationByID(idStr string) (*entities.Destination, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

	destination, err := service.Repo.DestinationByID(id)
	if err != nil {
		return nil, err
	}

	return destination, nil
}

func (service *DestinationService) DestinationsByLocationID(locationIDStr string) (*DestinationsByLocation, error) {
	var locationID uint
	if _, err := fmt.Sscanf(locationIDStr, "%d", &locationID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	location, err := service.LocationRepo.LocationByID(locationID)
	if err != nil {
		return &DestinationsByLocation{}, err
	}

	destinationIDs, err := service.Repo.DestinationIDsForLocation(locationID)
	var destinations []entities.Destination

	for _, destinationID := range destinationIDs {
		destination, err := service.Repo.DestinationByID(destinationID)
		if err != nil {
			return &DestinationsByLocation{}, err
		}
		destinations = append(destinations, *destination)
	}

	destinationsByLocation := &DestinationsByLocation{
		Location:     *location,
		Destinations: destinations,
	}

	return destinationsByLocation, nil
}

func (service *DestinationService) CreateDestination(destination entities.Destination) (entities.Destination, error) {
	_, err := service.LocationRepo.LocationByID(destination.LocationID)
	if err != nil {
		return entities.Destination{}, err
	}

	destination, err = service.Repo.CreateDestination(destination)
	if err != nil {
		return entities.Destination{}, err
	}
	return destination, nil
}

func (service *DestinationService) DeleteDestination(idStr string) (entities.Destination, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.Destination{}, errors.New("invalid ID format")
	}

	destination, err := service.Repo.DeleteDestination(id)
	if err != nil {
		return entities.Destination{}, err
	}
	return destination, nil
}

func (service *DestinationService) UpdateDestination(idStr string, destination entities.Destination) (entities.Destination, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.Destination{}, errors.New("invalid ID format")
	}

	destination, err := service.Repo.UpdateDestination(id, destination)
	if err != nil {
		return entities.Destination{}, err
	}
	return destination, nil
}

func (service *DestinationService) GenerateFakeLocation(f faker.Faker) (entities.Location, error) {
	fakeLocation := entities.Location{
		Name:        f.Address().City(),
		Country:     f.Address().Country(),
		Description: f.Lorem().Sentence(10),
	}

	location, err := service.LocationRepo.CreateLocation(fakeLocation)
	if err != nil {
		return entities.Location{}, err
	}

	return location, nil
}

func (service *DestinationService) GenerateFakeDestination(f faker.Faker) (entities.Destination, error) {
	min, max := 1, 9
	randomMultipleOfTen := f.IntBetween(min, max) * 10000

	location, err := service.GenerateFakeLocation(f)
	if err != nil {
		return entities.Destination{}, err
	}
	fakeDestination := entities.Destination{
		Name:             f.Company().Name(),
		LocationID:       location.ID,
		ImageUrl:         f.Internet().URL(),
		Description:      f.Lorem().Paragraph(3),
		VisitorsLastYear: randomMultipleOfTen,
		IsPrivate:        false,
	}

	destination, err := service.Repo.CreateDestination(fakeDestination)
	if err != nil {
		return entities.Destination{}, err
	}

	return destination, nil
}

func (service *DestinationService) StartGeneratingDestinations(interval time.Duration, f faker.Faker) {
	service.Ticker = time.NewTicker(interval)
	service.StopChan = make(chan bool)

	go func() {
		for {
			select {
			case <-service.Ticker.C:
				destination, err := service.GenerateFakeDestination(f)
				if err != nil {
					return
				}

				service.WsManager.AddToBroadcast(websocket.EventUpdateNotification{Action: "GenerateDestination", Destination: destination})

			case <-service.StopChan:
				return
			}
		}
	}()
}

func (service *DestinationService) StopGeneratingDestinations() {
	service.Ticker.Stop()
	//service.StopChan <- true
	return
}
