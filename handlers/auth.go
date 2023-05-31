package handlers

import (
	"net/http"

	"github.com/CodingCookieRookie/audit-log/handlers/ctrl"
	"github.com/gin-gonic/gin"
)

const (
	EMAIL_FIELD = "email"
	TOKEN_FIELD = "token"
)

func UserAuth(c *gin.Context) {
	email := c.GetHeader(EMAIL_FIELD)
	token := c.GetHeader(TOKEN_FIELD)

	if len(email) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, &ErrorMessage{
			Error: "email header missing in request",
		})
		return
	}

	if len(token) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, &ErrorMessage{
			Error: "token header missing in request",
		})
		return
	}

	err := ctrl.VerifyJWTToken(email, token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &ErrorMessage{
			Error: err.Error(),
		})

	}
	c.Next()
}
