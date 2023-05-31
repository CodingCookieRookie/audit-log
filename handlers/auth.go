package handlers

import (
	"net/http"

	"github.com/CodingCookieRookie/audit-log/handlers/ctrl"
	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	email := c.GetHeader("email")
	token := c.GetHeader("token")

	err := ctrl.VerifyJWTToken(email, token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &ErrorMessage{
			Error: err.Error(),
		})

	}
	c.Next()
}
