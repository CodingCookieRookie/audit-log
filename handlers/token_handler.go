package handlers

import (
	"fmt"

	"github.com/CodingCookieRookie/audit-log/constants"
	"github.com/CodingCookieRookie/audit-log/handlers/ctrl"
	"github.com/CodingCookieRookie/audit-log/log"
	"github.com/gin-gonic/gin"
)

type TokenGetRequest struct {
	Email string `form:"email"`
}

type TokenGetResponse struct {
	Message string `json:"message"`
}

func HandleApiTokenGet(c *gin.Context) (any, error) {
	if c.GetHeader("app-secret") != constants.GetAppSecret() {
		return nil, fmt.Errorf("incorrect app secret")
	}

	var tokenGetRequest TokenGetRequest

	if err := c.BindQuery(&tokenGetRequest); err != nil {
		log.Errorf("error binding token, err: %v", err)
		return nil, err
	}

	if err := ctrl.SendJWTTokenToEmail(tokenGetRequest.Email); err != nil {
		return nil, err
	}
	return &TokenGetResponse{
		Message: fmt.Sprintf("api token sent to %v successfully", tokenGetRequest.Email),
	}, nil
}
