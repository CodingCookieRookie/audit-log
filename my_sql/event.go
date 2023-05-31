package my_sql

import (
	"strings"

	"github.com/CodingCookieRookie/audit-log/models"
)

func createEventsTable() error {
	return Exec(`CREATE TABLE IF NOT EXISTS events (
		event_id int(11) NOT NULL AUTO_INCREMENT,
		event_type varchar(50) NOT NULL,
		event_time_stamp bigint(14),
		event_data_json varchar(6000),
		PRIMARY KEY(event_id), INDEX (event_type), INDEX(event_time_stamp)) 
	`)
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

func GetEvents(eventType string, startTimeStampMs, endTimeStampMs int) ([]*models.Event, error) {
	return Query(func(event *models.Event) []interface{} {
		return []interface{}{
			&event.EventType, &event.EventTimeStampMs, &event.EventDataJson,
		}
	},
		`SELECT
			event_type, event_time_stamp, event_data_json	
		FROM audit_log.events 
		WHERE
			event_type = ?
		AND 
			event_time_stamp <= ?
		AND 
			event_time_stamp >= ?`, eventType, endTimeStampMs, startTimeStampMs)
}

func GetEventByByTimeStamp(startTimeStampMs, endTimeStampMs int) ([]*models.Event, error) {
	return Query(func(event *models.Event) []interface{} {
		return []interface{}{
			&event.EventType, &event.EventTimeStampMs, &event.EventDataJson,
		}
	},
		`SELECT
			event_type, event_time_stamp, event_data_json	
		FROM audit_log.events 
		WHERE
			event_time_stamp <= ?
		AND 
			event_time_stamp >= ?`, endTimeStampMs, startTimeStampMs)
}
