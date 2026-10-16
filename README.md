# Job Scraper Platform

A monorepo for a job-search platform. The planned components are a Nuxt web app, a C# API, and a Go scraper.

## Repository layout

- `apps/web/` — Nuxt web application (planned)
- `services/api/` — C# API (planned)
- `services/scraper/` — Go scraper

Each component lives in its own directory so it can have its own dependencies, tools, and development workflow.

## Run the scraper

From `services/scraper`, run:

```sh
go run ./cmd/scraper "https://www.werkenbijns.nl/vacatures"
```

The scraper reads vacancy listings from the supplied page and follows its next-page links.
