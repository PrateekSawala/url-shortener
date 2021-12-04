package endpoint

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateShortURLSuccess(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:3002/createShortUrl/?url=https://google.com", nil)
	response := httptest.NewRecorder()

	CreateShortURLHandler(response, request)
	if response.Code != http.StatusOK {
		t.Error("response code is not 200")
	}
}
