package feeds2imap

import (
	readability "github.com/go-shiori/go-readability"
	"time"
)

// GetFullTextReadability fetches full article description for url
func GetFullTextReadability(url string) (string, error) {

	article, err := readability.FromURL(url, 30*time.Second)
	return article.Content, err

}
