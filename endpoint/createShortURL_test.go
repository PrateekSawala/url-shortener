package endpoint

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/PrateekSawala/url-shortener/domain"

	"github.com/stretchr/testify/assert"
)

func TestCreateShortURLSuccess(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:3002/createShortUrl/?url=https://google.com", nil)
	response := httptest.NewRecorder()

	CreateShortURLHandler(response, request)
	assert.Equal(t, response.Code, http.StatusOK)
}

func TestCreateShortURLError(t *testing.T) {

	t.Run("Should return error of empty input with response code 400", func(t *testing.T) {
		request := httptest.NewRequest("GET", "http://localhost:3002/createShortUrl/", nil)
		response := httptest.NewRecorder()

		CreateShortURLHandler(response, request)
		if response.Code != http.StatusBadRequest {
			t.Error("response code is not 400")
		}
		assert.EqualError(t, errors.New(strings.TrimSpace(response.Body.String())), domain.EmptyInput)
	})

	t.Run("Should return error of url is a already shorted with response code 400", func(t *testing.T) {
		request := httptest.NewRequest("GET", "http://localhost:3002/createShortUrl/?url=http://localhost:3002/shortUrl/2aa789c823074703b7baa8a524eb4aad", nil)
		response := httptest.NewRecorder()

		CreateShortURLHandler(response, request)
		if response.Code != http.StatusBadRequest {
			t.Error("response code is not 400")
		}
		assert.EqualError(t, errors.New(strings.TrimSpace(response.Body.String())), domain.URLIsAlreadyShortened)
	})

	t.Run("Should return error of invalid url with response code 400", func(t *testing.T) {
		request := httptest.NewRequest("GET", "http://localhost:3002/createShortUrl/?url=https://", nil)
		response := httptest.NewRecorder()

		CreateShortURLHandler(response, request)
		if response.Code != http.StatusBadRequest {
			t.Error("response code is not 400")
		}
		assert.EqualError(t, errors.New(strings.TrimSpace(response.Body.String())), domain.InvalidUrl)
	})
}
