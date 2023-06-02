package models

type Event struct {
	EventType        string `json:"event_type" binding:"required"`
	EventData        any    `json:"event_data"  binding:"required"` // specific event data
	EventTimeStampMs int    `json:"event_time_stamp_ms"`            // different from request timestamp
}
