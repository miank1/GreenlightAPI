package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If no route matches, ServeMux will return 404.
		// We intercept it here.
		_, pattern := mux.Handler(r)
		if pattern == "" {
			app.notFoundResponse(w, r)
			return
		}

		mux.ServeHTTP(w, r)
	})
}
