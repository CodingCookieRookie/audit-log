package handlers

import (
	"github.com/CodingCookieRookie/audit-log/handlers/ctrl"
	"github.com/gin-gonic/gin"
)

type UserPostRequest struct {
	Email  string `json:"email"`
	APIKey string `json:"api_key"`
}

type UserPostResponse struct {
}

// func HandleEventGet(c *gin.Context) (any, error) {
// 	var event Event

// 	if err := c.BindQuery(&event); err != nil {
// 		return nil, err
// 	}

// 	return &EventPostResponse{}, nil
// }

func HandleUserPost(c *gin.Context) (any, error) {
	var userPostRequest UserPostRequest

	if err := c.ShouldBindJSON(&userPostRequest); err != nil {
		return nil, err
	}

	err := ctrl.PostUser(userPostRequest.Email, userPostRequest.APIKey)

	return &EventPostResponse{}, err
}
