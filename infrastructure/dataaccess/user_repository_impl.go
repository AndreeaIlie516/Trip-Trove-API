package dataaccess

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
)

type GormUserRepository struct {
	Db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{Db: db}
}

func (r *GormUserRepository) AllUsers() ([]entities.User, error) {
	var users []entities.User
	result := r.Db.Find(&users)
	return users, result.Error
}

func (r *GormUserRepository) AllUserIDs() ([]uint, error) {
	var userIDs []uint

	if err := r.Db.Model(&entities.User{}).Select("ID").Find(&userIDs).Error; err != nil {
		return nil, err
	}

	return userIDs, nil
}

func (r *GormUserRepository) UserByID(id uint) (*entities.User, error) {
	var user entities.User

	if err := r.Db.First(&user, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *GormUserRepository) Register(user entities.User) (entities.User, error) {
	user.Role = entities.NormalUser

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entities.User{}, err
	}
	user.Password = string(hashedPassword)
	if err := r.Db.Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *GormUserRepository) Login(loginData entities.LoginRequest) (entities.LoginResponse, error) {
	var user entities.User

	if err := r.Db.First(&user, "email = ?", loginData.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.LoginResponse{}, errors.New("email not found")
		}
		return entities.LoginResponse{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return entities.LoginResponse{}, err
	}

	jwtWrapper := utils.JwtWrapper{
		SecretKey:         os.Getenv("JWT_SECRET"),
		Issuer:            "AuthService",
		ExpirationMinutes: 60,
	}

	signedToken, err := jwtWrapper.GenerateToken(user)
	if err != nil {
		return entities.LoginResponse{}, err
	}

	data := entities.LoginResponse{
		Email: user.Email,
		Jwt:   signedToken,
		ID:    user.ID,
	}

	return data, nil
}

func (r *GormUserRepository) DeleteUser(id uint) (entities.User, error) {
	var user entities.User

	if err := r.Db.First(&user, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, errors.New("user not found")
		}
		return entities.User{}, err
	}

	if err := r.Db.Delete(&user).Error; err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (r *GormUserRepository) UpdateUser(id uint, updatedUser entities.User) (entities.User, error) {
	var user entities.User

	if err := r.Db.First(&user, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, errors.New("user not found")
		}
		return entities.User{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return entities.User{}, err
	}
	updatedUser.Password = string(hashedPassword)

	if err := r.Db.Model(&user).Updates(updatedUser).Error; err != nil {
		return entities.User{}, err
	}

	return user, nil
}
