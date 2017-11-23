package main

import (
	"log"
	"os"

	"github.com/dstotijn/funda2csv"
)

func cli(fundaToken, searchOpts string) {
	r, err := funda.Search(fundaToken, searchOpts, 1, 25)
	if err != nil {
		log.Fatalf("Error: Could not search Funda: %v", err)
	}
	defer r.Close()

	objects, _, err := funda.ObjectsFromSearchResult(r)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Retrieved %d object(s).", len(objects))

	if err := funda.ObjectsToCSV(objects, os.Stdout); err != nil {
		log.Fatalf("Error: Could not write search result as CSV: %v", err)
	}
}
