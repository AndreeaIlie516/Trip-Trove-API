package middlewares

import (
	"Trip-Trove-API/domain/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"os"
	"strings"
)

type AuthMiddleware struct{}

func (rm AuthMiddleware) RequireRole(requiredRole entities.AccessType) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")
		if len(extractedToken) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect Format of Authorization Token"})
			c.Abort()
			return
		}
		clientToken = strings.TrimSpace(extractedToken[1])

		token, _ := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			secretKey := os.Getenv("JWT_SECRET")
			if secretKey == "" {
				log.Fatal("JWT_SECRET is not set in .env file")
			}

			return []byte(secretKey), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			roleFloat, ok := claims["role"].(float64)
			if !ok {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid role type in token"})
				c.Abort()
				return
			}

			role := entities.AccessType(roleFloat)

			if requiredRole != entities.NormalUser && role != requiredRole {
				c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
				c.Abort()
				return
			}
			c.Set("userID", claims["userID"])
			c.Set("role", role)

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}
