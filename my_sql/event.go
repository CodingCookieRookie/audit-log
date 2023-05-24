package my_sql

import (
	"strings"

	"github.com/CodingCookieRookie/audit-log/models"
)

func init() {

}
func InsertEvent(eventType string, eventTimestamp int, eventDataJson string) error {
	if len(eventType) == 0 {
		return nil
	}
	var query strings.Builder
	args := []any{eventType, eventTimestamp, eventDataJson}
	query.WriteString(returnPlaceHolderString(args))
	return Exec(
		`INSERT INTO audit_log.events (event_type, event_time_stamp, event_data_json) VALUES `+query.String(), args...)
}

func GetEventByType(eventType string) ([]*models.Event, error) {
	return Query(func(event *models.Event) []interface{} {
		return []interface{}{
			&event.EventType, &event.EventTimeStampMs, &event.EventData,
		}
	},
		`SELECT
			event_type, event_time_stamp, event_data	
		FROM audit_log.events 
		WHERE
			event_type = ?`, eventType)
}
