package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

type ListExecutator struct {
	service *calendar.Service
}

func NewListExecutator(service *calendar.Service) *ListExecutator {
	return &ListExecutator{
		service: service,
	}
}

func (le *ListExecutator) GetEvents() (*calendar.Events, error) {
	now := time.Now().Format(time.RFC3339)

	fmt.Println("Getting the upcoming 10 events")
	events, err := le.service.Events.List("primary").TimeMin(now).MaxResults(10).SingleEvents(true).OrderBy("startTime").Do()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve events: %v", err)
	}

	return events, nil
}

func (le *ListExecutator) Execute() {
	events, err := le.GetEvents()
	if err != nil {
		log.Fatalf("Failed to get events: %v", err)
	}

	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
		return
	}

	fmt.Println("Upcoming events:")
	for _, event := range events.Items {
		start := event.Start.DateTime
		if start == "" {
			start = event.Start.Date
		}
		fmt.Printf("%s - %s\n", start, event.Summary)
	}
}

func main() {
	listExecutator := NewListExecutator(service)
	listExecutator.Execute()
}
