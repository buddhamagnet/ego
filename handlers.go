package main

import (
	"encoding/json"
	"errors"
	"fmt"
	c "github.com/EconomistDigitalSolutions/ego/content"
	g "github.com/EconomistDigitalSolutions/ego/logging"
	"github.com/gorilla/mux"
	"net/http"
)

// Roothandler - root endpoint.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	json, _ := json.Marshal(RouteSet)
	w.Write(json)
}

// BlogsHandler - blog listings.
func ContentHandler(w http.ResponseWriter, r *http.Request) {
	// Get route parameters from the mux.
	vars := mux.Vars(r)
	// Assemble the blog listing url from the parameters and
	// the content source in the registry.
	if content, found := c.CONTENT_REGISTRY[vars["content"]]; !found {
		g.LogError(errors.New("resource not found"), "invalid-content-url", "content service api request errors")
		g.HttpError(w, 404, nil)
	} else {
		url := fmt.Sprintf(content["url"], vars["path"])
		if data, err := c.ProcessContent(url, vars["content"], w); err != nil {
			g.LogError(err, "application-fatal-error", "content service api fatal errors")
			g.HttpError(w, 500, err)
		} else {
			w.Write(data)
		}
	}
}
