package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(
	w http.ResponseWriter,
	status int,
	data interface{},
	headers http.Header,
) error {

	// Convert to JSON
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Add newline (cleaner output in browser)
	js = append(js, '\n')

	// Add extra headers if provided
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Write status
	w.WriteHeader(status)

	// Write response
	w.Write(js)

	return nil
}

// Generic error response
func (app *application) errorResponse(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	message interface{},
) {
	env := map[string]interface{}{
		"error": message,
	}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// 500 Internal Server Error
func (app *application) serverErrorResponse(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	app.logger.Println(err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// 404 Not Found
func (app *application) notFoundResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}
