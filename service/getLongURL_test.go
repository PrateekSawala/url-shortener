package service

import (
	"url-shortener/domain"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLongUrlSuccess(t *testing.T) {
	response, err := GetLongURL("2aa789c823074703b7baa8a524eb4aad")
	if err != nil {
		t.Errorf("GetLongURL Error: %s", err)
		return
	}
	result := &domain.UrlInfo{Url: "https://google.com"}
	assert.Equal(t, result, response)
	assert.NoError(t, err)
}
