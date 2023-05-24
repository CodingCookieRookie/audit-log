package handlers

import (
	"github.com/CodingCookieRookie/audit-log/handlers/ctrl"
	"github.com/CodingCookieRookie/audit-log/models"
	"github.com/gin-gonic/gin"
)

type EventPostResponse struct {
	PostStatus string        `json:"post_status"`
	Event      *models.Event `json:"event"`
}

// func HandleEventGet(c *gin.Context) (any, error) {
// 	var event Event

// 	if err := c.BindQuery(&event); err != nil {
// 		return nil, err
// 	}

// 	return &EventPostResponse{}, nil
// }

func HandleEventPost(c *gin.Context) (any, error) {
	var event *models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		return nil, err
	}

	err := ctrl.PostEvent(event)

	return &EventPostResponse{}, err
}
