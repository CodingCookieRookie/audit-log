package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CodingCookieRookie/audit-log/handlers"
	"github.com/CodingCookieRookie/audit-log/models"
	"github.com/CodingCookieRookie/audit-log/my_sql"
	"github.com/gin-gonic/gin"
)

type EventHandler interface {
	HandleEvent() error
}

// Event represents a recorded event with common and specific fields

// In-memory storage for simplicity (not suitable for production)
var eventList []EventHandler

// Handler for event submission
func submitEvent(w http.ResponseWriter, r *http.Request) {
	// Verify authentication here

	// Parse request body into an Event object
	var event models.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to parse event: %v", err)
		return
	}

	// Set the timestamp for the event
	//event.EventTimeStampMs = time.Now()

	// Store the event in the in-memory storage
	//eventList = append(eventList, event)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Event recorded successfully")
}

// Handler for querying events by field values
func queryEvents(w http.ResponseWriter, r *http.Request) {
	// Verify authentication here

	// Parse query parameters
	//fieldValue := r.URL.Query().Get("field")
	// Add more query parameters as needed

	// Search for events matching the field value
	var matchingEvents []models.Event

	for _, event := range eventList {
		// Perform matching logic based on the query parameters
		// For example:
		// if event.Fields.SomeField == fieldValue {
		//     matchingEvents = append(matchingEvents, event)
		// }
		event.HandleEvent()
		// Add more matching conditions as needed
	}

	// Serialize the matching events as JSON and send the response
	response, err := json.Marshal(matchingEvents)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to serialize events: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// func main() {
// 	// Create a new router
// 	r := mux.NewRouter()

// 	// Define HTTP endpoints
// 	r.HandleFunc("/events", submitEvent).Methods("POST")
// 	r.HandleFunc("/events/query", queryEvents).Methods("GET")

// 	// Start the HTTP server on port 8080
// 	log.Println("Audit log service started on port 8080")
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

func main() {
	engine := gin.Default()
	my_sql.Init()
	handlers.Route(engine)
	engine.Run(":3000")
}
