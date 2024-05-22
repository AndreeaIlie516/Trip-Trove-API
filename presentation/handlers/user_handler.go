package handlers

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/services"
	"Trip-Trove-API/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserHandler struct {
	Service *services.UserService
}

func (handler *UserHandler) AllUsers(c *gin.Context) {
	users, err := handler.Service.AllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (handler *UserHandler) UserByID(c *gin.Context) {
	requestedID := c.Param("id")
	userIDInterface, _ := c.Get("userID")
	role, _ := c.Get("role")

	userIDFloat, _ := userIDInterface.(float64)

	var reqID uint
	_, err := fmt.Sscan(requestedID, &reqID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if role == entities.NormalUser && uint(userIDFloat) != reqID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	user, err := handler.Service.UserByID(requestedID)
	if err != nil {
		fmt.Println("Error fetching user:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (handler *UserHandler) Register(c *gin.Context) {
	var newUser entities.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	validators := map[string]validator.Func{
		"usernameValidator": utils.UsernameValidator,
		"nameValidator":     utils.NameValidator,
		"passwordValidator": utils.PasswordValidator,
	}

	print("Aici")

	for validatorName, validatorFunction := range validators {
		if err := validate.RegisterValidation(validatorName, validatorFunction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register validator: " + validatorName})
			return
		}
	}

	err := validate.Struct(newUser)

	if err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid validation error"})
			return
		}

		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := "Validation error on field '" + err.Field() + "': " + err.ActualTag()
			if err.Param() != "" {
				errorMessage += " (Parameter: " + err.Param() + ")"
			}
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	user, err := handler.Service.Register(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (handler *UserHandler) Login(c *gin.Context) {
	var loginData entities.LoginRequest

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponse, err := handler.Service.Login(loginData)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "login failed"})
		return
	}

	c.JSON(http.StatusOK, loginResponse)
}

func (handler *UserHandler) DeleteUser(c *gin.Context) {
	requestedID := c.Param("id")
	userIDInterface, _ := c.Get("userID")
	role, _ := c.Get("role")

	userIDFloat, _ := userIDInterface.(float64)

	var reqID uint
	_, err := fmt.Sscan(requestedID, &reqID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if role == entities.NormalUser && uint(userIDFloat) != reqID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	user, err := handler.Service.DeleteUser(requestedID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (handler *UserHandler) UpdateUser(c *gin.Context) {
	requestedID := c.Param("id")
	userIDInterface, _ := c.Get("userID")
	role, _ := c.Get("role")

	userIDFloat, _ := userIDInterface.(float64)

	var reqID uint
	_, err := fmt.Sscan(requestedID, &reqID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if role == entities.NormalUser && uint(userIDFloat) != reqID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	var updatedUser entities.User

	if err := c.BindJSON(&updatedUser); err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	validate := validator.New()

	validators := map[string]validator.Func{
		"usernameValidator": utils.UsernameValidator,
		"nameValidator":     utils.NameValidator,
		"passwordValidator": utils.PasswordValidator,
	}

	for validatorName, validatorFunction := range validators {
		if err := validate.RegisterValidation(validatorName, validatorFunction); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register validator: " + validatorName})
			return
		}
	}

	err = validate.Struct(updatedUser)

	if err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid validation error"})
			return
		}

		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := "Validation error on field '" + err.Field() + "': " + err.ActualTag()
			if err.Param() != "" {
				errorMessage += " (Parameter: " + err.Param() + ")"
			}
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	user, err := handler.Service.UpdateUser(requestedID, updatedUser)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
