package handlers

import (
	"github.com/CodingCookieRookie/audit-log/handlers/ctrl"
	"github.com/CodingCookieRookie/audit-log/log"
	"github.com/gin-gonic/gin"
)

type TokenGetRequest struct {
	Email string `json:"email"`
}

type TokenGetResponse struct {
	Message string `json:"message"`
}

func HandleApiToken(c *gin.Context) (any, error) {
	var tokenGetRequest TokenGetRequest

	if err := c.BindHeader(&tokenGetRequest); err != nil {
		log.Errorf("error binding token, err: %v", err)
		return nil, err
	}

	if err := ctrl.SendJWTTokenToEmail(tokenGetRequest.Email); err != nil {
		return nil, err
	}
	return &TokenGetResponse{
		Message: "Token posted successfully",
	}, nil
}
