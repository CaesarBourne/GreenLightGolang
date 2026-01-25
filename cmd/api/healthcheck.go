package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)

	jsonResponse := `{ "status": "available", "environment": %q, "version": %q}`
	jsonResponse = fmt.Sprintf(jsonResponse, app.config.env, version)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(jsonResponse))
}
