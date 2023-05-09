package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

type InsertExecutator struct {
	service *calendar.Service
}

func NewInsertExecutator(service *calendar.Service) *InsertExecutator {
	return &InsertExecutator{
		service: service,
	}
}

func (ie *InsertExecutator) Execute() {
	requestID := uuid.New().String()

	event := &calendar.Event{
		Summary:     "Google I/O 2015",
		Location:    "800 Howard St., San Francisco, CA 94103",
		Description: "A chance to hear more about Google's developer products.",
		Start: &calendar.EventDateTime{
			DateTime: "2023-05-08T10:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		End: &calendar.EventDateTime{
			DateTime: "2023-05-08T11:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		Recurrence: []string{
			"RRULE:FREQ=DAILY;COUNT=2",
		},
		Attendees: []*calendar.EventAttendee{
			{Email: "jose.vsilva@mercadolivre.com"},
		},
		Reminders: &calendar.EventReminders{
			UseDefault: false,
			Overrides: []*calendar.EventReminder{
				{Method: "email", Minutes: 24 * 60},
				{Method: "popup", Minutes: 10},
			},
		},
		ConferenceData: &calendar.ConferenceData{
			CreateRequest: &calendar.CreateConferenceRequest{
				RequestID: requestID,
				ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
					Type: "hangoutsMeet",
				},
			},
		},
	}

	event, err := ie.service.Events.Insert("primary", event).ConferenceDataVersion(1).Do()
	if err != nil {
		log.Fatalf("Failed to create event: %v", err)
	}

	fmt.Printf("Event created: %s\n", event.HtmlLink)
	fmt.Printf("Event: %+v\n", event)
}

func main() {
	insertExecutator := NewInsertExecutator(service)
	insertExecutator.Execute()
}
