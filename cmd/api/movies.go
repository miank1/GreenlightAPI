package main

import (
	"net/http"
)

type Movie struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Year    int      `json:"year"`
	Runtime int      `json:"runtime"`
	Genres  []string `json:"genres"`
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		app.errorResponse(w, r, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var input struct {
		Title   string   `json:"title"`
		Year    int      `json:"year"`
		Runtime int      `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	// Decode JSON
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "invalid JSON body")
		return
	}

	v := NewValidator()

	if input.Title == "" {
		v.AddError("title", "must be provided")
	}

	if input.Year <= 1888 {
		v.AddError("year", "must be greater than 1888")
	}

	if input.Runtime <= 0 {
		v.AddError("runtime", "must be positive")
	}

	if len(input.Genres) == 0 {
		v.AddError("genres", "must contain at least one value")
	}

	if !v.Valid() {
		app.errorResponse(w, r, http.StatusUnprocessableEntity, v.Errors)
		return
	}

	// Simulate saving (no DB yet)
	movie := Movie{
		ID:      1,
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	data := map[string]interface{}{
		"movie": movie,
	}

	err = app.writeJSON(w, http.StatusCreated, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
