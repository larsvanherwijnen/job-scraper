package scraper

import (
	"errors"
	"net/url"
	"strings"
)

// Vacancy contains the information extracted from one vacancy listing.
type Vacancy struct {
	Title       string
	Location    string
	Hours       string
	Discipline  string
	Description string
	URL         string
}

// Scrape fetches the starting listing page and follows its next-page links.
func Scrape(startURL string) ([]Vacancy, error) {
	if err := validateURL(startURL); err != nil {
		return nil, err
	}

	var vacancies []Vacancy
	currentURL := startURL
	visitedURLs := make(map[string]bool)

	for currentURL != "" && !visitedURLs[currentURL] {
		visitedURLs[currentURL] = true

		doc, err := fetchDocument(currentURL)
		if err != nil {
			return nil, err
		}

		baseURL, err := url.Parse(currentURL)
		if err != nil {
			return nil, err
		}

		vacancies = append(vacancies, parseVacancies(doc, baseURL)...)

		nextURL, exists, err := findNextPageURL(doc, baseURL)
		if err != nil {
			return nil, err
		}
		if !exists {
			break
		}
		currentURL = nextURL
	}

	return vacancies, nil
}

func validateURL(rawURL string) error {
	parsedURL, err := url.Parse(rawURL)
	if err != nil ||
		(parsedURL.Scheme != "http" && parsedURL.Scheme != "https") ||
		parsedURL.Hostname() == "" ||
		!strings.Contains(parsedURL.Hostname(), ".") {
		return errors.New("invalid URL")
	}
	return nil
}
