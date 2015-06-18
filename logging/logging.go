package logging

import (
	log "github.com/Sirupsen/logrus"
	"github.com/Sirupsen/logrus/hooks/syslog"
	"log/syslog"
	"net/http"
	"os"
)

var guids = map[string]string{
	"invalid-content-url":     "9425a36f-57a7-4754-b6f9-b1aa7b1780a2",
	"application-fatal-error": "ecdcaf55-81c8-465f-86a2-fce38ea1ba67",
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	hook, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_DEBUG, "")
	if err != nil {
		log.Error("Unable to connect to local syslog daemon")
	} else {
		log.AddHook(hook)
	}
}

// Log writes the error to the log stream.
func LogError(err error, event, topic string) {
	if err != nil {
		log.WithFields(log.Fields{
			"event": event,
			"topic": topic,
			"key":   getMessage(event),
		}).Error(err)
	}
}

// HttpError outputs the appropriate HTTP error
// response from the given status code.
func HttpError(w http.ResponseWriter, code int, err error) {
	http.Error(w, http.StatusText(code), code)
}

func getMessage(message string) string {
	return guids[message]
}
