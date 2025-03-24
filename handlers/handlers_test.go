package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreateUserHandler(t *testing.T) {
	// Clean up the test environment
	defer func() {
		_ = os.Remove("user.json")
	}()

	t.Run("No body supplied", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/user/create", nil)
		rec := httptest.NewRecorder()

		CreateUserHandler(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
		}

		expected := "no body supplied\n"
		if rec.Body.String() != expected {
			t.Errorf("expected body %q, got %q", expected, rec.Body.String())
		}
	})

	t.Run("Create user when file does not exist", func(t *testing.T) {
		user := User{
			FirstName:    "John",
			LastName:     "Doe",
			EmailAddress: "john.doe@example.com",
		}
		body, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/user/create", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		CreateUserHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
		}

		data, err := ioutil.ReadFile("user.json")
		if err != nil {
			t.Fatalf("expected file to be created, but got error: %v", err)
		}

		var users []User
		err = json.Unmarshal(data, &users)
		if err != nil {
			t.Fatalf("failed to unmarshal users: %v", err)
		}

		if len(users) != 1 || users[0] != user {
			t.Errorf("expected users to contain %v, got %v", user, users)
		}
	})

	t.Run("Append user when file exists", func(t *testing.T) {
		initialUser := User{
			FirstName:    "Jane",
			LastName:     "Smith",
			EmailAddress: "jane.smith@example.com",
		}
		initialData, _ := json.Marshal([]User{initialUser})
		_ = ioutil.WriteFile("user.json", initialData, 0777)

		newUser := User{
			FirstName:    "Alice",
			LastName:     "Johnson",
			EmailAddress: "alice.johnson@example.com",
		}
		body, _ := json.Marshal(newUser)
		req := httptest.NewRequest(http.MethodPost, "/user/create", bytes.NewReader(body))
		rec := httptest.NewRecorder()

		CreateUserHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
		}

		data, err := ioutil.ReadFile("user.json")
		if err != nil {
			t.Fatalf("expected file to exist, but got error: %v", err)
		}

		var users []User
		err = json.Unmarshal(data, &users)
		if err != nil {
			t.Fatalf("failed to unmarshal users: %v", err)
		}

		if len(users) != 2 || users[0] != initialUser || users[1] != newUser {
			t.Errorf("expected users to contain %v and %v, got %v", initialUser, newUser, users)
		}
	})
}
