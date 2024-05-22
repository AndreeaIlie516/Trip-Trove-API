package dataaccess

import (
	"Trip-Trove-API/domain/entities"
	"errors"
	"gorm.io/gorm"
)

type GormDestinationRepository struct {
	Db *gorm.DB
}

func NewGormDestinationRepository(db *gorm.DB) *GormDestinationRepository {
	return &GormDestinationRepository{Db: db}
}

func (r *GormDestinationRepository) AllDestinations() ([]entities.Destination, error) {
	var destinations []entities.Destination
	result := r.Db.Find(&destinations)
	return destinations, result.Error
}

func (r *GormDestinationRepository) AllDestinationIDs() ([]uint, error) {
	var destinationIDs []uint

	if err := r.Db.Model(&entities.Destination{}).Select("ID").Find(&destinationIDs).Error; err != nil {
		return nil, err
	}

	return destinationIDs, nil
}

func (r *GormDestinationRepository) DestinationByID(id uint) (*entities.Destination, error) {
	var destination entities.Destination

	if err := r.Db.First(&destination, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("destination not found")
		}
		return nil, err
	}

	return &destination, nil
}

func (r *GormDestinationRepository) DestinationIDsForLocation(locationID uint) ([]uint, error) {
	var destinationIDs []uint

	if err := r.Db.Where("location_id = ?", locationID).Model(&entities.Destination{}).Select("ID").Find(&destinationIDs).Error; err != nil {
		return nil, err
	}

	return destinationIDs, nil
}

func (r *GormDestinationRepository) DeleteDestinationsByLocationID(locationID uint) error {
	if err := r.Db.Where("location_id = ?", locationID).Delete(&entities.Destination{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormDestinationRepository) CreateDestination(destination entities.Destination) (entities.Destination, error) {
	if err := r.Db.Create(&destination).Error; err != nil {
		return entities.Destination{}, err
	}
	return destination, nil
}

func (r *GormDestinationRepository) DeleteDestination(id uint) (entities.Destination, error) {
	var destination entities.Destination

	if err := r.Db.First(&destination, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Destination{}, errors.New("destination not found")
		}
		return entities.Destination{}, err
	}

	if err := r.Db.Delete(&destination).Error; err != nil {
		return entities.Destination{}, err
	}

	return destination, nil
}

func (r *GormDestinationRepository) UpdateDestination(id uint, updatedDestination entities.Destination) (entities.Destination, error) {
	var destination entities.Destination

	if err := r.Db.First(&destination, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Destination{}, errors.New("destination not found")
		}
		return entities.Destination{}, err
	}

	if err := r.Db.Model(&destination).Updates(updatedDestination).Error; err != nil {
		return entities.Destination{}, err
	}

	return destination, nil
}
