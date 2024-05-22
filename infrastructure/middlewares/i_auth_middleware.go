package middlewares

import (
	"Trip-Trove-API/domain/entities"
	"github.com/gin-gonic/gin"
)

type IAuthMiddleware interface {
	RequireRole(role entities.AccessType) gin.HandlerFunc
}
