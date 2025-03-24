package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type User struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
}

type Users struct {
	Users []User
}

func RegisterRoutes(r *chi.Mux) {
	r.Route("/user", func(r chi.Router) {
		r.Post("/create", CreateUserHandler)
		r.Get("/listusers", ListUsersHandler)
	})
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "no body supplied", http.StatusBadRequest)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body", http.StatusInternalServerError)
		return
	}

	var user User
	err = json.Unmarshal(bodyBytes, &user)

	if _, err := os.Stat("user.json"); errors.Is(err, os.ErrNotExist) {
		var users Users
		users.Users = append(users.Users, user)

		usersBytes, err := json.Marshal(users.Users)
		if err != nil {
			http.Error(w, "unable to marshal user", http.StatusInternalServerError)
			return
		}
		err = os.WriteFile("user.json", usersBytes, 0777)
		if err != nil {
			http.Error(w, "unable to write file", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	readFile, err := os.ReadFile("user.json")
	if err != nil {
		http.Error(w, "unable to read file", http.StatusInternalServerError)
		return
	}

	var us Users
	err = json.Unmarshal(readFile, &us.Users)
	if err != nil {
		http.Error(w, "unable to unmarshal updated users", http.StatusInternalServerError)
		return
	}

	us.Users = append(us.Users, user)

	updatedUsersBytes, err := json.Marshal(us.Users)
	if err != nil {
		http.Error(w, "unable to marshal updated users", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("user.json", updatedUsersBytes, 0777)
	if err != nil {
		http.Error(w, "unable to write updated users to file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	userBytes, err := os.ReadFile("user.json")
	if err != nil {
		http.Error(w, "unable to process users", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(userBytes)
	if err != nil {
		http.Error(w, "unable to process users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
