package endpoint

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
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
	if response.Code != http.StatusSeeOther {
		t.Error("response code is not 302")
	}
}
