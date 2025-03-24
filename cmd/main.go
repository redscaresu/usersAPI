package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/redscaresu/natwest/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	handlers.RegisterRoutes(r)
	http.ListenAndServe(":8080", r)
}
