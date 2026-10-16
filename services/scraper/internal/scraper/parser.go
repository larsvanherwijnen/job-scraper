package scraper

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parseVacancies(doc *goquery.Document, baseURL *url.URL) []Vacancy {
	var vacancies []Vacancy

	doc.Find("ul.vacancy-list > li.vacancy-item-cell > a.vacancy-item").Each(func(_ int, item *goquery.Selection) {
		title := strings.TrimSpace(item.Find(".vacancy-item-titles h3").First().Text())
		if title == "" {
			return
		}

		href, exists := item.Attr("href")
		if !exists {
			return
		}

		jobURL, err := url.Parse(href)
		if err != nil {
			return
		}

		vacancies = append(vacancies, Vacancy{
			Title:       title,
			Location:    strings.TrimSpace(item.Find("li.location .metadata-text").First().Text()),
			Hours:       strings.TrimSpace(item.Find("li.hours .metadata-text").First().Text()),
			Discipline:  strings.TrimSpace(item.Find("li.discipline .metadata-text").First().Text()),
			Description: strings.TrimSpace(item.Find(".vacancy-item-body > .text").First().Text()),
			URL:         baseURL.ResolveReference(jobURL).String(),
		})
	})

	return vacancies
}

func findNextPageURL(doc *goquery.Document, baseURL *url.URL) (string, bool, error) {
	nextHref, exists := doc.Find(".pagination.arrow-right a").Attr("href")
	if !exists {
		return "", false, nil
	}

	nextURL, err := url.Parse(nextHref)
	if err != nil {
		return "", false, err
	}

	return baseURL.ResolveReference(nextURL).String(), true, nil
}
