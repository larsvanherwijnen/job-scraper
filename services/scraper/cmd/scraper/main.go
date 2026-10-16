package main

import (
	"fmt"
	"os"

	"job-scraper/services/scraper/internal/scraper"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <vacancy-list-url>")
		return
	}

	vacancies, err := scraper.Scrape(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, vacancy := range vacancies {
		fmt.Printf("Titel: %s\nLocatie: %s\nUren: %s\nVakgebied: %s\nOmschrijving: %s\nURL: %s\n\n",
			vacancy.Title,
			vacancy.Location,
			vacancy.Hours,
			vacancy.Discipline,
			vacancy.Description,
			vacancy.URL,
		)
	}

	fmt.Println("Totaal:", len(vacancies))
}
