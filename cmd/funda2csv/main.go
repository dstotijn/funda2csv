package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	searchOpts := flag.String("searchOpts", "/amsterdam/1-dag", "Funda search options")
	httpServer := flag.Bool("http", false, "Run as HTTP server")
	addr := flag.String("addr", ":8080", "Address for HTTP server to listen on")
	flag.Parse()

	fundaToken := os.Getenv("FUNDA_ALERT_FUNDA_TOKEN")
	if fundaToken == "" {
		log.Fatal("Error: Environment variable `FUNDA_ALERT_FUNDA_TOKEN` cannot be empty.")
	}

	if *httpServer {
		http.Handle("/funda.csv", handler(fundaToken))
		log.Fatal(http.ListenAndServe(*addr, nil))
	}

	cli(fundaToken, *searchOpts)
}
