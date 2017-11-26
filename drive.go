// Shamelessly adapted from https://developers.google.com/drive/v3/web/quickstart/go 20171125
package main

import (
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"

	drive "google.golang.org/api/drive/v3"
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

	svc, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve drive Client %v", err)
	}

	return svc
}

/*
	drive := driveService()

		r, err := drive.Files.List().Q("name contains 'astrolabe'").PageSize(10).
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
*/
