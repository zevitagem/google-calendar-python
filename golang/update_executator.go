package main

import (
	"fmt"
	"log"

	"google.golang.org/api/calendar/v3"
)

type UpdateExecutator struct {
	service *calendar.Service
	id      string
}

func NewUpdateExecutator(service *calendar.Service) *UpdateExecutator {
	return &UpdateExecutator{
		service: service,
	}
}

func (ue *UpdateExecutator) Configure(config map[string]string) {
	ue.id = config["id"]
}

func (ue *UpdateExecutator) Execute() {
	eventService := ue.service.Events

	event, err := eventService.Get("primary", ue.id).Do()
	if err != nil {
		log.Fatalf("Failed to retrieve event: %v", err)
	}

	event.Summary = "Appointment at Somewhere - Update"

	updatedEvent, err := eventService.Update("primary", ue.id, event).Do()
	if err != nil {
		log.Fatalf("Failed to update event: %v", err)
	}

	fmt.Printf("Event updated: %s\n", updatedEvent.Updated)
}

func main() {
	updateExecutator := NewUpdateExecutator(service)

	updateExecutator.Configure(map[string]string{"id": "827gj35869mlimobovmajc1m94"})
	updateExecutator.Execute()
}
