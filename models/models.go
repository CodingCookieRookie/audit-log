package models

type Event struct {
	EventType        string `json:"event_type"`
	EventData        any    `json:"event_data"`          // specific event data
	EventTimeStampMs int    `json:"event_time_stamp_ms"` // different from request timestamp
}

// type UserAccount struct { // for new and deactivate account event data
// 	Username        string  `json:"username"`
// 	UserEmail       string  `json:"user_email"`
// 	CreatorUsername string  `json:"creator_username"`
// 	CreatorEmail    string  `json:"creator_email"` // creator
// 	Country         string  `json:"country"`
// 	Amount          float64 `json:"amount"`
// 	Currency        string  `json:"currency"`
// }

// type PasswordChange struct {
// 	Username  string `json:"username"`
// 	IPAddress string `json:"ip_address"`
// }

// type Bill struct { // for bill event data
// 	BillAmount    float64 `json:"bill_amount"`
// 	Currency      string  `json:"currency"`
// 	CurrentAmount float64 `json:"current_amount"`
// 	BillDetails   string  `json:"bill_details"`
// }
