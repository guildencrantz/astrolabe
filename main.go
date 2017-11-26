package main

import (
	"log"

	"github.com/spf13/viper"
)

func main() {
	log.Print("start")
	defer func() { log.Print("finish") }()
	// TODO: Need two modes: train and run.
	run()
}

func run() {
	sheets := []TargetSheet{}
	if err := viper.UnmarshalKey(SHEETS_KEY, &sheets); err != nil {
		log.Fatalf("Unable to get sheets.\n%e", err)
	}

	sht = sheetsService()

	for _, s := range sheets {
		ss, err := sht.Spreadsheets.Get(s.ID).IncludeGridData(true).Do()
		if err != nil {
			log.Fatalf("Unable to get sheet '%s': %e", s.ID, err)
		}

		spread := &Spreadsheet{Spreadsheet: ss}
		spread.SetTransactionSheet()
		t := spread.transactions.Transactions()

		log.Printf("%#q", t)
	}
}
