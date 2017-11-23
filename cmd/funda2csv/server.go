package main

import (
	"log"
	"net/http"

	funda "github.com/dstotijn/funda2csv"
)

func handler(fundaToken string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := funda.Search(fundaToken, "/amsterdam/1-dag", 1, 25)
		if err != nil {
			log.Printf("Error: Could not search Funda: %v", err)
			internalServerError(w)
			return
		}
		defer res.Close()

		objects, _, err := funda.ObjectsFromSearchResult(res)
		if err != nil {
			log.Print(err)
			internalServerError(w)
			return
		}
		log.Printf("Retrieved %d object(s).", len(objects))

		w.Header().Set("Content-Type", "text/csv; charset=utf-8")

		if err := funda.ObjectsToCSV(objects, w); err != nil {
			log.Printf("Error: Could not write search result as CSV: %v", err)
			internalServerError(w)
			return
		}
	})
}

func internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}
