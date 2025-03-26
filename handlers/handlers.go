package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/redscaresu/usersAPI/handlers/models"
)

type application struct {
	filesystem string
}

func NewApplication(filesystem string) *application {
	return &application{
		filesystem: filesystem,
	}
}

func (a *application) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("/create", a.CreateUserHandler)
	r.HandleFunc("/listusers", a.ListUsersHandler)
	r.HandleFunc("/updateuser", a.UpdateUserHandler)

}

func (a *application) Create(user *models.User) error {

	err := os.Chdir(a.filesystem)
	if err != nil {
		return errors.New("unable to cd to dir")
	}

	if _, err := os.Stat("user.json"); errors.Is(err, os.ErrNotExist) {
		var users models.Users
		users.Users = append(users.Users, *user)

		usersBytes, err := json.Marshal(users.Users)
		if err != nil {
			return errors.New("unable to marshal users")
		}
		err = os.WriteFile("user.json", usersBytes, 0777)
		if err != nil {
			return errors.New("unable to write file")
		}
		return nil
	}

	readFile, err := os.ReadFile("user.json")
	if err != nil {
		return errors.New("unable to read file")
	}

	var us models.Users
	err = json.Unmarshal(readFile, &us.Users)
	if err != nil {
		return errors.New("unable to unmarshal updated users")
	}

	us.Users = append(us.Users, *user)
	updatedUsersBytes, err := json.Marshal(us.Users)
	if err != nil {
		return err
	}

	err = os.WriteFile("user.json", updatedUsersBytes, 0777)
	if err != nil {
		return errors.New("unable to write users file")

	}

	return nil
}

func (a *application) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "no body supplied", http.StatusBadRequest)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body", http.StatusInternalServerError)
		return
	}

	var user models.User
	err = json.Unmarshal(bodyBytes, &user)

	err = a.Create(&user)
	if err != nil {
		http.Error(w, "unable to marshal user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func (a *application) ListUsersHandler(w http.ResponseWriter, r *http.Request) {

	//check the error
	os.Chdir(a.filesystem)
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

	return
}

func (a *application) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "no body supplied", http.StatusBadRequest)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body", http.StatusInternalServerError)
		return
	}

	var userUpdate models.User
	err = json.Unmarshal(bodyBytes, &userUpdate)
	if err != nil {
		http.Error(w, "unable to process user update", http.StatusInternalServerError)
		return
	}

	userFile, err := os.ReadFile("user.json")
	if err != nil {
		http.Error(w, "unable to read file", http.StatusInternalServerError)
		return
	}

	var currentUsers models.Users
	err = json.Unmarshal(userFile, &currentUsers.Users)
	if err != nil {
		http.Error(w, "unable to process current users", http.StatusInternalServerError)
		return
	}

	for i, v := range currentUsers.Users {
		if v.FirstName == userUpdate.FirstName {
			currentUsers.Users[i] = userUpdate
		}
	}

	currentUserByte, err := json.Marshal(currentUsers.Users)
	if err != nil {
		http.Error(w, "unable to process current users", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("user.json", currentUserByte, 0777)
	if err != nil {
		http.Error(w, "unable to write updated users to file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
