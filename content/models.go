package content

// Whitehouse progress feed.
type Whitehouse struct {
	Category string `json:"category"`
	UID      int    `json:"uid"`
	Body     string `json:"body"`
	URLTitle string `json:"url_title"`
	Path     string `json:"path"`
}
