package main

import (
	"fmt"
	"net/http"
	"time"

	"caesar.emmanuel.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The movie is created successfully")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	movieId, err := app.readIDParams(r)

	if err != nil || movieId < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "this is the movie id boss %d\n", movieId)

	movieData := data.Movie{ID: movieId, CreatedAt: time.Now(), Title: "Emmanuel",
		Runtime: 102, Genres: []string{"thriller", "crime", " romance"}}

	err = app.writeJSON(w, http.StatusOK, movieData, nil)

	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The ser er has an error", http.StatusInternalServerError)
	}

}
