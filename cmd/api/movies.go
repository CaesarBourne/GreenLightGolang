package main

import (
	"fmt"
	"net/http"
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
}
