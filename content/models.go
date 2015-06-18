package content

// Generic type containing values
// common to all content.
type Content struct {
	Title   string `json:"title"`
	Body    string `json:"body"`
	Teaser  string `json:"teaser"`
	Status  string `json:"status"`
	Created string `json:"created"`
}

// Economist blog post.
type Article struct {
	Title   string `json:"title"`
	Body    string `json:"body"`
	Teaser  string `json:"teaser"`
	Status  string `json:"status"`
	Created string `json:"created"`
	Comment string `json:"comment"`
}

// Whitehouse progress feed.
type Whitehouse struct {
	Category string `json:"category"`
	UID      int    `json:"uid"`
	Body     string `json:"body"`
	URLTitle string `json:"url_title"`
	Path     string `json:"path"`
}
