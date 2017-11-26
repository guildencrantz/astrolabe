// Shamelessly adapted from https://developers.google.com/drive/v3/web/quickstart/go 20171125
package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"

	drive "google.golang.org/api/drive/v3"
)

var (
	drv *drive.Service
)

func driveService() *drive.Service {
	b, err := ioutil.ReadFile(confFilePath(OAUTH_CLIENT_SECRET_KEY))
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	drv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve drive Client %v", err)
	}

	return drv
}

func QueryFiles(q string) {
	// q := "name contains 'astrolabe'"
	r, err := drv.Files.List().Q(q).PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	log.Println("Files:")
	if len(r.Files) > 0 {
		for _, i := range r.Files {
			log.Printf("%s (%s)\n", i.Name, i.Id)
		}
	} else {
		fmt.Println("No files found.")
	}
}
