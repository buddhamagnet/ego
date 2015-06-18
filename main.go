package main

import (
	"flag"
	"github.com/EconomistDigitalSolutions/ego/content"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
)

var port string

func init() {
	// Set port to listen on, default to 9494.
	flag.StringVar(&port, "port", ":9494", "port to listen on")
	if port == "" {
		port = ":9494"
	}
	// The environment variables are autoloaded via
	// the silent godotenv import above. The service will
	// complain if critical variables such as the data
	// source are not set when we attempt to retrieve content
	// which will happen when we call...
	content.Setup()
}

func main() {
	flag.Parse()
	NewRouter()
	log.Println("ego tripping on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
