package main

import (
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	sheets "google.golang.org/api/sheets/v4"
)

func sheetsService() *sheets.Service {
	b, err := ioutil.ReadFile(confFilePath(OAUTH_CLIENT_SECRET_KEY))
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	config, err := google.ConfigFromJSON(b, sheets.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	svc, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve sheets Client %v", err)
	}

	return svc
}
