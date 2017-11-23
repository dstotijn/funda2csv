package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
)

func main() {
	searchOpts := flag.String("searchOpts", "/amsterdam/1-dag", "Funda search options")
	flag.Parse()

	fundaToken := os.Getenv("FUNDA_ALERT_FUNDA_TOKEN")
	if fundaToken == "" {
		log.Fatal("Error: Environment variable `FUNDA_ALERT_FUNDA_TOKEN` cannot be empty.")
	}

	r, err := searchFunda(fundaToken, *searchOpts, 1, 25)
	if err != nil {
		log.Fatalf("Error: Could not search Funda: %v", err)
	}
	defer r.Close()

	objects, _, err := fundaObjectsFromSearchResult(r)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Retrieved %d object(s).", len(objects))

	w := csv.NewWriter(os.Stdout)
	w.Write([]string{"id", "address", "url"})
	for _, o := range objects {
		if err := w.Write([]string{
			o.id,
			o.address,
			o.url.String(),
		}); err != nil {
			log.Fatalf("Error: Could not write CSV record: %v", err)
		}
	}
	w.Flush()
}
