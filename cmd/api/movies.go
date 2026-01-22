package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The movie is created successfully")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	movieId, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	if err != nil || movieId < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "this is the movie id boss %d\n", movieId)
}
