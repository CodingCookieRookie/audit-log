package handlers

import (
	"fmt"
	"strconv"
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
	EventTimeStampGMT   string `form:"gmt"`
	EventOrder          string `form:"event_order"`
}

type EventPostResponse struct {
	Status string        `json:"status"`
	Event  *models.Event `json:"event"`
}

func parseGMT(gmt string) (int, error) {

	invalidGMTLength := len(gmt) > 3 || len(gmt) < 2
	invalidGMTOperator := gmt[0:1] != "-" && gmt[0:1] != "*"

	if invalidGMTLength || invalidGMTOperator {
		log.Errorf("invalid gmt length: %v, invalid gmt operator: %v", invalidGMTLength, invalidGMTOperator)
		return 0, fmt.Errorf(constants.GMT_INVALID)
	}

	gmtHours, err := strconv.Atoi(gmt[1:])
	if err != nil {
		log.Errorf("error converting gmt to numerical value, err: %v", err)
		return 0, fmt.Errorf(constants.GMT_INVALID)
	}

	if gmtHours < 0 || gmtHours > 12 {
		log.Errorf("error converting gmt to numerical value, err: %v", err)
		return 0, fmt.Errorf(constants.GMT_OUT_OF_RANGE)
	}

	totalGMTHoursInMs := gmtHours * int(time.Hour.Milliseconds())
	log.Infof("totalGMTHoursInMs: %v", totalGMTHoursInMs)

	if gmt[0] == '*' {
		return -totalGMTHoursInMs, nil
	}

	return totalGMTHoursInMs, nil
}

func HandleEventGet(c *gin.Context) (any, error) {
	var eventGetRequest EventGetRequest

	if err := c.BindQuery(&eventGetRequest); err != nil {
		return nil, err
	}

	userEmail := c.GetHeader(EMAIL_FIELD)

	var endTimeStampMs int
	var startTimeStampMs int
	var gmtMs int

	if len(eventGetRequest.EventTimeStampGMT) != 0 {
		var err error
		gmtMs, err = parseGMT(eventGetRequest.EventTimeStampGMT)
		if err != nil {
			return nil, err
		}
	}

	log.Infof("gmtMs: %v", gmtMs)

	if len(eventGetRequest.EventTimeStampEnd) != 0 {
		timeStampEnd, err := time.Parse(constants.TIME_FORMAT, eventGetRequest.EventTimeStampEnd)
		if err != nil {
			log.Errorf("error parsing event time stamp end, err: %v", err)
			return nil, fmt.Errorf("error parsing event time stamp end, err: %v", err)
		} else {
			endTimeStampMs = int(timeStampEnd.UnixMilli()) + gmtMs
		}

	}

	if endTimeStampMs == 0 {
		endTimeStampMs = int(time.Now().UnixMilli()) // default time stamp end to time now
	}

	if len(eventGetRequest.EventTimeStampStart) != 0 {
		timeStampStart, err := time.Parse(constants.TIME_FORMAT, eventGetRequest.EventTimeStampStart)
		if err != nil {
			log.Errorf("error parsing event time stamp start, err: %v", err)
			return nil, fmt.Errorf("error parsing event time stamp start, err: %v", err)
		} else {
			startTimeStampMs = int(timeStampStart.UnixMilli()) + gmtMs
		}

	}

	log.Infof("startTimeStampMs: %v", startTimeStampMs)
	log.Infof("endTimeStampMs: %v", endTimeStampMs)

	if startTimeStampMs > endTimeStampMs {
		return nil, fmt.Errorf("start time stamp ms value invalid")
	}

	invalidEventOrder := eventGetRequest.EventOrder != constants.EVENT_ORDER_ASC && eventGetRequest.EventOrder != constants.EVENT_ORDER_DESC

	if len(eventGetRequest.EventOrder) == 0 { // default event order to desc
		eventGetRequest.EventOrder = constants.EVENT_ORDER_DESC
	} else if invalidEventOrder {
		return nil, fmt.Errorf("event order is invalid, order is only ASC or DESC")
	}

	events, err := ctrl.GetEvents(userEmail, eventGetRequest.EventType, startTimeStampMs, endTimeStampMs, eventGetRequest.EventOrder)

	if err != nil {
		log.Errorf("error getting events, err: %v", err)
		return nil, err
	}

	for _, event := range events {
		eventDataJson := event.EventData
		event.EventData = string(eventDataJson.([]byte))
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
