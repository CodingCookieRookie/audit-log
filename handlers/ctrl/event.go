package ctrl

import (
	"encoding/json"

	"github.com/CodingCookieRookie/audit-log/log"
	"github.com/CodingCookieRookie/audit-log/models"
	"github.com/CodingCookieRookie/audit-log/my_sql"
)

func PostEvent(event *models.Event) error {
	eventDataJson, err := json.Marshal(event.EventData)
	if err != nil {
		log.Errorf("error marshalling event data: %v", err)
		return err
	}

	return my_sql.InsertEvent(event.EventType, event.EventTimeStampMs, string(eventDataJson))
}
