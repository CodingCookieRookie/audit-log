package ctrl

import (
	"encoding/json"

	"github.com/CodingCookieRookie/audit-log/log"
	"github.com/CodingCookieRookie/audit-log/models"
	"github.com/CodingCookieRookie/audit-log/my_sql"
)

func PostEvent(event *models.Event) error {
	eventDataJson, err := json.Marshal(event.EventDataJson)
	if err != nil {
		log.Errorf("error marshalling event data: %v", err)
		return err
	}

	return my_sql.InsertEvent(event.EventType, event.EventTimeStampMs, string(eventDataJson))
}

func GetEvents(eventType string, startTimeStampMs, endTimeStampMS int) ([]*models.Event, error) {
	if len(eventType) == 0 {
		return my_sql.GetEventByByTimeStamp(startTimeStampMs, endTimeStampMS)
	}
	return my_sql.GetEvents(eventType, startTimeStampMs, endTimeStampMS)
}
