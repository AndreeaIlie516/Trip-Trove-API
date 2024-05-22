package dataaccess

import (
	"Trip-Trove-API/domain/entities"
	"errors"
	"gorm.io/gorm"
)

type GormLocationRepository struct {
	Db *gorm.DB
}

func NewGormLocationRepository(db *gorm.DB) *GormLocationRepository {
	return &GormLocationRepository{Db: db}
}

func (r *GormLocationRepository) AllLocations() ([]entities.Location, error) {
	var locations []entities.Location
	result := r.Db.Find(&locations)
	return locations, result.Error
}

func (r *GormLocationRepository) AllLocationIDs() ([]uint, error) {
	var locationIDs []uint

	if err := r.Db.Model(&entities.Location{}).Select("ID").Find(&locationIDs).Error; err != nil {
		return nil, err
	}

	return locationIDs, nil
}

func (r *GormLocationRepository) LocationByID(id uint) (*entities.Location, error) {
	var location entities.Location

	if err := r.Db.First(&location, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("location not found")
		}
		return nil, err
	}

	return &location, nil
}

func (r *GormLocationRepository) CreateLocation(location entities.Location) (entities.Location, error) {
	if err := r.Db.Create(&location).Error; err != nil {
		return entities.Location{}, err
	}
	return location, nil
}

func (r *GormLocationRepository) DeleteLocation(id uint) (entities.Location, error) {
	var location entities.Location

	if err := r.Db.First(&location, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Location{}, errors.New("location not found")
		}
		return entities.Location{}, err
	}

	if err := r.Db.Delete(&location).Error; err != nil {
		return entities.Location{}, err
	}

	return location, nil
}

func (r *GormLocationRepository) UpdateLocation(id uint, updatedLocation entities.Location) (entities.Location, error) {
	var location entities.Location

	if err := r.Db.First(&location, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Location{}, errors.New("location not found")
		}
		return entities.Location{}, err
	}

	if err := r.Db.Model(&location).Updates(updatedLocation).Error; err != nil {
		return entities.Location{}, err
	}

	return location, nil
}
