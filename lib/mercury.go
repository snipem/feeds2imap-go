package feeds2imap

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

type mercuryResponse struct {
	Author        string      `json:"author"`
	Content       string      `json:"content"`
	DatePublished interface{} `json:"date_published"`
	Dek           interface{} `json:"dek"`
	Direction     string      `json:"direction"`
	Domain        string      `json:"domain"`
	Excerpt       string      `json:"excerpt"`
	LeadImageURL  string      `json:"lead_image_url"`
	NextPageURL   interface{} `json:"next_page_url"`
	RenderedPages int64       `json:"rendered_pages"`
	Title         string      `json:"title"`
	TotalPages    int64       `json:"total_pages"`
	URL           string      `json:"url"`
	WordCount     int64       `json:"word_count"`
}

// GetFullText fetches full article description for url
func GetFullText(url string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://mercury.postlight.com/parser?url="+url, nil)
	if err != nil {
		return "", err
	}

	apiKey := viper.GetString("mercury.apiKey")

	req.Header.Add("x-api-key", apiKey)
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	response := mercuryResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	return response.Content, nil

}
