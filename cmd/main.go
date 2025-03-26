package main

import (
	"net/http"

	"github.com/redscaresu/usersAPI/handlers"
)

func main() {
	r := http.NewServeMux()

	application := handlers.NewApplication(".")

	application.RegisterRoutes(r)
	http.ListenAndServe(":8080", r)
}
