package scraper

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func fetchDocument(pageURL string) (*goquery.Document, error) {
	response, err := http.Get(pageURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, fmt.Errorf("unexpected HTTP status: %s", response.Status)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
