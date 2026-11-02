# Groupie Tracker

A Go web application that collects data about music artists from a public REST API and presents it as a clean, browsable website with band members, first albums, concert locations and tour dates.

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?logo=css3&logoColor=white)
![REST API](https://img.shields.io/badge/REST-API-informational)

## Overview

The server fetches four related API resources and combines them into a single view for each artist:

| Endpoint | Data |
|---|---|
| `/api/artists` | Name, image, members, creation year, first album |
| `/api/locations` | Cities where the artist has performed |
| `/api/dates` | Concert dates |
| `/api/relation` | Which dates belong to which location |

The project uses only the Go standard library (`net/http`, `encoding/json`, `html/template`).

## Features

- **Artist gallery**: a responsive grid of artist cards with hover animations
- **Artist detail page** (`/artists/{id}`) showing:
  - Image, creation date and first album
  - All band members
  - Concert locations and dates, linked through the relation data
- **Error handling**: custom pages for `400` (non-numeric ID), `404` (unknown route or artist) and `500` (API or template failure)
- Dark theme with glassmorphism cards and smooth transitions

## Getting Started

**Requirements:** Go 1.20 or newer

```bash
git clone https://github.com/ahmed-jawher/groupie-tracker.git
cd groupie-tracker
go run .
```

Open <http://localhost:8080>.

> **Note:** The project was built against the course API at `groupietrackers.herokuapp.com`, which is no longer online. To run it today, point the URLs in `handlers.go` to a compatible mirror or a local mock of the four endpoints.

## Routes

| Route | Description |
|---|---|
| `GET /` | Artist gallery |
| `GET /artists/{id}` | Details for a single artist |

## Project Structure

```text
groupie-tracker/
├── main.go                            # server setup and routing
├── handlers.go                        # HTTP handlers, API calls, error pages
├── models.go                          # API and view models
├── templates/
│   ├── index.html                     # artist gallery
│   ├── artist.html                    # artist details
│   └── error.html                     # shared error page
└── groupie-tracker-visualizations/    # follow-up project (see below)
```

## Related Projects

| Project | Description |
|---|---|
| [`groupie-tracker-visualizations/`](groupie-tracker-visualizations) | The follow-up stage, focused on presenting the data in a clear, visual way |

## Author

**Ahmed Jawher** — [@ahmed-jawher](https://github.com/ahmed-jawher)
