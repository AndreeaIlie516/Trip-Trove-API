package utils

import (
	"Trip-Trove-API/domain/entities"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtWrapper struct {
	SecretKey         string
	Issuer            string
	ExpirationMinutes int64
	ExpirationHours   int64
}

type JwtClaim struct {
	Email  string              `json:"email"`
	Role   entities.AccessType `json:"role"`
	UserID uint                `json:"userID"`
	jwt.StandardClaims
}

func (j *JwtWrapper) GenerateToken(user entities.User) (signedToken string, err error) {
	claims := &JwtClaim{
		Email:  user.Email,
		Role:   user.Role,
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(j.ExpirationMinutes) * time.Minute).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (j *JwtWrapper) RefreshToken(user entities.User) (signedToken string, err error) {
	claims := &JwtClaim{
		Email:  user.Email,
		Role:   user.Role,
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}
	return
}

func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return
	}
	return
}
