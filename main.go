package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	// TODO: Need two modes: train and run.
	run()
}

func run() {

	sheets := []*sheet{}
	if err := viper.UnmarshalKey(SHEETS_KEY, &sheets); err != nil {
		log.Fatal("Unable to get sheets.", err)
	}

	svc := sheetsService()

	for _, s := range sheets {
		f, err := svc.Spreadsheets.Get(s.Id).Do()
		if err != nil {
			log.Fatalf("Unable to get sheet '%s': %e", s.Id, err)
		}

		log.Printf("Retrieved '%s'", f.SpreadsheetUrl)
	}
}
