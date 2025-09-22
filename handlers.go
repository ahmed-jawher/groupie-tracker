package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

// ErrorPage represents the data for the error template
type ErrorPage struct {
	Code    int
	Title   string
	Message string
}

func handleError(w http.ResponseWriter, status int) {
	var page ErrorPage

	switch status {
	case http.StatusNotFound:

		page = ErrorPage{404, "Page Not Found", "Oops! The page you are looking for has gone on tour! 🎸"}
	case http.StatusBadRequest:
		page = ErrorPage{400, "Bad Request", "Oops! Something's wrong with this request! 🤔"}
	case http.StatusInternalServerError:
		page = ErrorPage{500, "Server Error", "Our backstage crew is working on fixing this! 🛠️"}
	default:
		page = ErrorPage{status, "Error", "Something went wrong!"}
	}

	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, page.Title, status)
		log.Printf("Error loading error template: %v", err)
		return
	}

	w.WriteHeader(status)
	t.Execute(w, page)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleError(w, http.StatusNotFound)
		return
	}

	data := getArtists()
	if data == nil {
		handleError(w, http.StatusInternalServerError)
		return
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		handleError(w, http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	// Extracting the ID from the URL
	id := path.Base(r.URL.Path[len("/artists/"):])

	// if id is not a number, return 400 Bad Request
	if _, err := strconv.Atoi(id); err != nil {
		handleError(w, http.StatusBadRequest)
		return
	}

	// Fetch the artist by ID
	data := getArtistByID(id)
	if data == nil {
		handleError(w, http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		handleError(w, http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
