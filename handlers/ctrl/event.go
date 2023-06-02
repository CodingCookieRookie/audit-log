package ctrl

import (
	"encoding/json"

	"github.com/CodingCookieRookie/audit-log/log"
	"github.com/CodingCookieRookie/audit-log/models"
	"github.com/CodingCookieRookie/audit-log/my_sql"
)

func PostEvent(userEmail string, event *models.Event) error {
	eventDataJson, err := json.Marshal(event.EventData)
	if err != nil {
		log.Errorf("error marshalling event data: %v", err)
		return err
	}

	return my_sql.InsertEvent(userEmail, event.EventType, event.EventTimeStampMs, string(eventDataJson))
}

func GetEvents(userEmail, eventType string, startTimeStampMs, endTimeStampMs int, eventOrder string) ([]*models.Event, error) {
	if len(eventType) == 0 {
		return my_sql.GetEventByByTimeStamp(userEmail, startTimeStampMs, endTimeStampMs, eventOrder)
	}
	return my_sql.GetEvents(userEmail, eventType, startTimeStampMs, endTimeStampMs, eventOrder)
}
