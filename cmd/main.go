package main

import (
	"net/http"

	"github.com/redscaresu/usersAPI/handlers"
)

func main() {
	r := http.NewServeMux()

	handlers.RegisterRoutes(r)
	http.ListenAndServe(":8080", r)
}
