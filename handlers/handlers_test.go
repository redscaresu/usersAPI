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
	"github.com/stretchr/testify/require"
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

	buf, err := json.Marshal(expectedUser)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewReader(buf))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	t.Cleanup(func() {
		os.Remove("user.json")
	})
}
