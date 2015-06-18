package content

import (
	"encoding/json"
	"errors"
	e "github.com/EconomistDigitalSolutions/ego/logging"
	"io/ioutil"
	"net/http"
)

// ProcessContent retrieves JSON from a content source, transforms
// it and returns it in a normalised format to the consumer.
func ProcessContent(url, content string, w http.ResponseWriter) ([]byte, error) {
	resp, err := http.Get(url)
	// If we get a response back from the content source, make sure we
	// get a 200 OK.
	if resp != nil {
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			e.HttpError(w, resp.StatusCode, errors.New("ouch!"))
			return nil, nil
		}
	}
	// If we get an error from the content source, log it and return a 500 to the client.
	// This includes network errors, errors processing and transforming the JSON and any
	// other issues cleanly processing the request. The error will be captured by the
	// handler and return a 500 to the client to prevent the service from crashing.
	if err != nil {
		e.LogError(err, "content-service-data-source-error", "content service data source")
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	// Issue processing the JSON from the content source.
	if err != nil {
		e.LogError(err, "content-service-data-read-error", "content service data processing")
		return nil, err
	}
	j, err := TransformContent(content, data)
	// Issue transforming the content for consumption by the front end.
	if err != nil {
		e.LogError(err, "content-service-data-transform-error", "content service data processing")
		return nil, err
	}
	return j, nil
}

// TransformContent takes an interface, decodes the original JSON,
// marshals it into the value and sends the transformed JSON back
// down the pipe.
func TransformContent(content string, data []byte) ([]byte, error) {
	template := MODEL_REGISTRY[content]
	err := json.Unmarshal(data, &template)

	if err != nil {
		return nil, err
	}
	data, err = json.Marshal(&template)
	if err != nil {
		return nil, err
	}
	return data, nil
}
