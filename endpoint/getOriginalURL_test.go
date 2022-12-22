package endpoint

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PrateekSawala/url-shortener/domain"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetLongURLSuccess(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:3002/shortUrl", nil)
	response := httptest.NewRecorder()

	// Set gorilla/mux vars
	vars := map[string]string{
		"id": "2aa789c823074703b7baa8a524eb4aad",
	}
	request = mux.SetURLVars(request, vars)

	GetOriginalURLHandler(response, request)
	assert.Equal(t, response.Code, http.StatusSeeOther)
}

func TestGetLongURLError(t *testing.T) {

	t.Run("Should return error of empty input with response code 400", func(t *testing.T) {
		request := httptest.NewRequest("GET", "http://localhost:3002/shortUrl", nil)
		response := httptest.NewRecorder()

		GetOriginalURLHandler(response, request)
		if response.Code != http.StatusBadRequest {
			t.Error("response code is not 400")
		}
		assert.EqualError(t, errors.New(strings.TrimSpace(response.Body.String())), domain.EmptyInput)
	})

	t.Run("Should return error of invalid Input with response code 400", func(t *testing.T) {
		request := httptest.NewRequest("GET", "http://localhost:3002/shortUrl", nil)
		response := httptest.NewRecorder()
		// Set gorilla/mux vars
		vars := map[string]string{
			"id": "123456",
		}
		request = mux.SetURLVars(request, vars)

		GetOriginalURLHandler(response, request)
		if response.Code != http.StatusBadRequest {
			t.Error("response code is not 400")
		}
		assert.EqualError(t, errors.New(strings.TrimSpace(response.Body.String())), domain.InvalidInput)
	})

	t.Run("Should return error of url not found with response code 500", func(t *testing.T) {
		request := httptest.NewRequest("GET", "http://localhost:3002/shortUrl", nil)
		response := httptest.NewRecorder()

		// Set gorilla/mux vars
		vars := map[string]string{
			"id": "2aa789c823074703b7baa8a524eb4aab",
		}
		request = mux.SetURLVars(request, vars)

		GetOriginalURLHandler(response, request)
		if response.Code != http.StatusInternalServerError {
			t.Error("response code is not 500")
		}
		assert.EqualError(t, errors.New(strings.TrimSpace(response.Body.String())), domain.URLNotFound)
	})
}
