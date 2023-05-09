package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func auth() (*oauth2.Token, error) {
	ctx := context.Background()

	// If modifying these scopes, delete the file token.json.
	scopes := []string{calendar.CalendarScope}

	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokenFile := "./token.json"
	credentialsFile := "./credentials.json"

	creds, err := google.FindDefaultCredentials(ctx, scopes...)
	if err != nil || creds == nil {
		creds, err = google.LegacyTokenSource(ctx, scopes...).Token()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve OAuth token: %v", err)
		}
	} else if creds.HasExpired() {
		if creds.RefreshToken != "" {
			err = creds.Refresh(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to refresh OAuth token: %v", err)
			}
		} else {
			return nil, fmt.Errorf("OAuth token has expired and no refresh token is available")
		}
	}

	// Save the credentials for the next run
	tokenData, err := json.Marshal(creds.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal token data: %v", err)
	}
	err = os.WriteFile(tokenFile, tokenData, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to write token file: %v", err)
	}

	return creds.Token, nil
}

func main() {
	ctx := context.Background()

	credsToken, err := auth()
	if err != nil {
		log.Fatalf("Failed to authenticate: %v", err)
	}

	credentialsFile := "./credentials.json"
	b, err := os.ReadFile(credentialsFile)
	if err != nil {
		log.Fatalf("Failed to read credentials file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Failed to parse credentials file: %v", err)
	}

	client := config.Client(ctx, credsToken)

	service, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Failed to create Calendar service: %v", err)
	}

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalf("Please provide an action (list, insert, or update)")
	}

	action := args[0]

	switch action {
	case "list":
		executator = &ListExecutator{Service: service} // Assuming ListExecutator, InsertExecutator, and UpdateExecutator are defined
	case "insert":
		executator = &InsertExecutator{Service: service}
	case "update":
		if len(args) < 2 {
			log.Fatalf("Please provide an event ID for update")
		}
		executator = &UpdateExecutator{
			Service: service,
			Config: map[string]string{
				"id": args[1],
			},
		}
	default:
		log.Fatalf("Invalid action: %s", action)
	}

	err = executator.Execute()
	if err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
}
