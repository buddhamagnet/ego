package content

var CONTENT_REGISTRY map[string]map[string]string
var MODEL_REGISTRY map[string]interface{}

func Setup() {
	CONTENT_REGISTRY = map[string]map[string]string{
		"whitehouse": map[string]string{
			"description": "Whitehouse progress",
			"url":         "https://www.whitehouse.gov/facts/json/progress/%s",
		},
	}

	MODEL_REGISTRY = map[string]interface{}{
		"whitehouse": []Whitehouse{},
	}
}
