package content

import (
	"os"
)

var CONTENT_REGISTRY map[string]map[string]string
var MODEL_REGISTRY map[string]interface{}

func Setup() {
	CONTENT_REGISTRY = map[string]map[string]string{
		"blogs": map[string]string{
			"description": "Economist blogs listing",
			"url":         os.Getenv("CONTENT_SERVICE_BASE_URL") + "/blogs/%s/contentasjson",
		},
		"article": map[string]string{
			"description": "Economist article",
			"url":         os.Getenv("CONTENT_SERVICE_BASE_URL") + "/news/%s/contentasjson",
		},
		"daily": map[string]string{
			"description": "Espresso daily content",
			"url":         os.Getenv("ESPRESSO_BASE_URL") + "/%s",
		},
		"whitehouse": map[string]string{
			"description": "Whitehouse progress",
			"url":         "https://www.whitehouse.gov/facts/json/progress/%s",
		},
	}

	MODEL_REGISTRY = map[string]interface{}{
		"article":    &Article{},
		"blogs":      &Content{},
		"daily":      &Content{},
		"whitehouse": []Whitehouse{},
	}
}
