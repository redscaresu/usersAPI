package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/redscaresu/usersAPI/handlers"
	"github.com/redscaresu/usersAPI/handlers/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	r := http.NewServeMux()
	application := handlers.NewApplication("test_directory")
	application.RegisterRoutes(r)

	expectedUser := &models.User{
		FirstName:    "lewis",
		LastName:     "Jones",
		EmailAddress: "foo@foo.com",
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(expectedUser)

	req := httptest.NewRequest(http.MethodPost, "/create", &buf)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	t.Cleanup(func() {
		os.Remove("user.json")
	})
}
