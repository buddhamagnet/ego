package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("ego front end running on port 9595")
	http.ListenAndServe(":9595", nil)
}
