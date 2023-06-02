package my_sql

import (
	"fmt"
	"strings"

	"github.com/CodingCookieRookie/audit-log/models"
)

func createEventsTable() error {
	return Exec(`CREATE TABLE IF NOT EXISTS events (
		event_id int(11) NOT NULL AUTO_INCREMENT,
		event_type varchar(50) NOT NULL,
		event_time_stamp bigint(20) NOT NULL,
		event_data_json varchar(6000),
		user_email varchar(320) NOT NULL,
		PRIMARY KEY(event_id), INDEX(user_email, event_type, event_time_stamp), INDEX(user_email, event_time_stamp), INDEX(event_time_stamp))
	`)
}

func InsertEvent(userEmail, eventType string, eventTimestamp int, eventDataJson string) error {
	if len(eventType) == 0 {
		return nil
	}
	var query strings.Builder
	args := []any{userEmail, eventType, eventTimestamp, eventDataJson}
	query.WriteString(returnPlaceHolderString(args))
	return Exec(
		fmt.Sprintf(`INSERT INTO %v.events (user_email, event_type, event_time_stamp, event_data_json) VALUES `, dbName)+query.String(), args...)
}

func GetEvents(userEmail, eventType string, startTimeStampMs, endTimeStampMs int, eventOrder string) ([]*models.Event, error) {
	return Query(func(event *models.Event) []interface{} {
		return []interface{}{
			&event.EventType, &event.EventTimeStampMs, &event.EventData,
		}
	},
		fmt.Sprintf(`SELECT
			event_type, event_time_stamp, event_data_json	
		FROM %v.events 
		WHERE
			user_email = ?
		AND
			event_type = ?
		AND 
			event_time_stamp <= ?
		AND 
			event_time_stamp >= ?
		ORDER BY
			event_time_stamp %v`, dbName, eventOrder), userEmail, eventType, endTimeStampMs, startTimeStampMs)
}

func GetEventByByTimeStamp(userEmail string, startTimeStampMs, endTimeStampMs int, eventOrder string) ([]*models.Event, error) {
	return Query(func(event *models.Event) []interface{} {
		return []interface{}{
			&event.EventType, &event.EventTimeStampMs, &event.EventData,
		}
	},
		fmt.Sprintf(`SELECT
			event_type,	event_time_stamp, event_data_json	
		FROM %v.events 
		WHERE
			user_email = ?
		AND
			event_time_stamp <= ?
		AND 
			event_time_stamp >= ?
		ORDER BY
			event_time_stamp %v`, dbName, eventOrder), userEmail, endTimeStampMs, startTimeStampMs)
}
