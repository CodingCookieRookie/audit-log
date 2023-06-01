package handlers

import (
	"fmt"
	"time"

	"github.com/CodingCookieRookie/audit-log/constants"
	"github.com/CodingCookieRookie/audit-log/handlers/ctrl"
	"github.com/CodingCookieRookie/audit-log/log"
	"github.com/CodingCookieRookie/audit-log/models"
	"github.com/gin-gonic/gin"
)

type EventGetRequest struct {
	EventType           string `form:"event_type"`
	EventTimeStampEnd   string `form:"event_timestamp_end"`   // Date in string with format 2006-01-02 15:04:05
	EventTimeStampStart string `form:"event_timestamp_start"` // Date in string with format 2006-01-02 15:04:05
}

type EventPostResponse struct {
	Status string        `json:"status"`
	Event  *models.Event `json:"event"`
}

func HandleEventGet(c *gin.Context) (any, error) {
	var eventGetRequest EventGetRequest

	if err := c.BindQuery(&eventGetRequest); err != nil {
		return nil, err
	}

	userEmail := c.GetHeader(EMAIL_FIELD)

	var endTimeStampMS int
	var startTimeStampMs int

	if len(eventGetRequest.EventTimeStampEnd) != 0 {
		timeStampEnd, err := time.Parse(constants.TIME_FORMAT, eventGetRequest.EventTimeStampEnd)
		if err != nil {
			log.Errorf("error parsing event time stamp end, err: %v", err)
			return nil, fmt.Errorf("error parsing event time stamp end, err: %v", err)
		} else {
			endTimeStampMS = int(timeStampEnd.UnixMilli())
		}

	}

	if endTimeStampMS == 0 {
		endTimeStampMS = int(time.Now().UnixMilli()) // default time stamp end to time now
	}

	if len(eventGetRequest.EventTimeStampStart) != 0 {
		timeStampStart, err := time.Parse(constants.TIME_FORMAT, eventGetRequest.EventTimeStampStart)
		if err != nil {
			log.Errorf("error parsing event time stamp start, err: %v", err)
			return nil, fmt.Errorf("error parsing event time stamp start, err: %v", err)
		} else {
			startTimeStampMs = int(timeStampStart.UnixMilli())
		}

	}

	log.Infof("endTimeStampMS: %v", endTimeStampMS)
	log.Infof("startTimeStampMs: %v", startTimeStampMs)

	if startTimeStampMs > endTimeStampMS {
		return nil, fmt.Errorf("start time stamp ms value invalid")
	}

	events, err := ctrl.GetEvents(userEmail, eventGetRequest.EventType, startTimeStampMs, endTimeStampMS)

	if err != nil {
		log.Errorf("error getting events, err: %v", err)
		return nil, err
	}

	for _, event := range events {
		eventDataJson := event.EventDataJson
		event.EventDataJson = string(eventDataJson.([]byte))
	}

	return events, err
}

func HandleEventPost(c *gin.Context) (any, error) {
	var event *models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		return nil, err
	}

	event.EventTimeStampMs = int(time.Now().UnixMilli())

	email := c.GetHeader(EMAIL_FIELD)

	err := ctrl.PostEvent(email, event)

	eventPostResponse := &EventPostResponse{
		Status: "post event to database success",
		Event:  event,
	}

	if err != nil {
		eventPostResponse.Status = "post event to database failed"
	}

	return eventPostResponse, err
}
