package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)

	// jsonResponse := `{ "status": "available", "environment": %q, "version": %q}`
	// jsonResponse = fmt.Sprintf(jsonResponse, app.config.env, version)

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	js, err := json.Marshal(data)

	if err != nil {
		app.logger.Print(err)
		http.Error(w, "error in server", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	// w.Write([]byte(jsonResponse))
	w.Write(js)
}
