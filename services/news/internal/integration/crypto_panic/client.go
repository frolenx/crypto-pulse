package crypto_panic

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/crypto-pulse/news/internal/model"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func FetchNews() ([]*model.News, error) {
	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		return nil, errors.New("API_TOKEN environment variable not set")
	}

	authToken := buildUrl(apiToken)

	resp, err := http.Get(authToken)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var getNewsResponse *model.GetNewsResponse
	err = json.Unmarshal(body, &getNewsResponse)
	if err != nil {
		return nil, err
	}

	var filtered []*model.News
	for _, news := range getNewsResponse.Results {
		t, err := time.Parse("2006-01-02T15:04:05Z", news.PublishedAt)
		if err != nil {
			return nil, err
		}

		if time.Since(t) < 30*time.Minute {
			desc, err := extractDescription(news.Url)
			if err != nil {
				return nil, err
			}
			news.Description = desc

			filtered = append(filtered, news)
		}
	}

	return filtered, nil
}

func buildUrl(authToken string) string {
	apiUrl := url.URL{
		Scheme: "https",
		Host:   "cryptopanic.com",
	}

	apiUrl.Path = "/api/v1/posts/"

	queryParams := apiUrl.Query()
	queryParams.Set("access_token", authToken)
	queryParams.Set("public", "true")
	apiUrl.RawQuery = queryParams.Encode()

	return apiUrl.String()
}

func extractDescription(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	description, exists := document.Find("meta[property='og:description']").Attr("content")
	if !exists {
		return "", errors.New("description not found")
	}

	return description, nil
}
