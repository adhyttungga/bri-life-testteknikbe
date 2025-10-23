package middleware

import (
	"net/http"
	"strings"

	"github.com/adhyttungga/bri-life-testteknikbe/models/dto"
	"github.com/adhyttungga/bri-life-testteknikbe/utils"
	"github.com/gin-gonic/gin"
)

func ProtectRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		resError := dto.ResponseError{
			ResponseCode: "01",
			ResponseDesc: http.StatusText(http.StatusUnauthorized),
		}

		// Get authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, resError)
			return
		}

		// Validate access token
		var agentId string
		accessToken := strings.TrimPrefix(authHeader, "Bearer ")
		if !utils.ValidateToken(accessToken, &agentId) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, resError)
			return
		}

		c.Set("agent_id", agentId)
		return
	}
}
