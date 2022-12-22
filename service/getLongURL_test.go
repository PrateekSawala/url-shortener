package service

import (
	"testing"

	"github.com/PrateekSawala/url-shortener/domain"

	"github.com/stretchr/testify/assert"
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

func TestGetLongUrlError(t *testing.T) {
	t.Run("Should return error of invalid Input", func(t *testing.T) {
		_, err := GetLongURL("")
		assert.EqualError(t, domain.ErrInvalidInput, err.Error())
	})

	t.Run("Should return error of url not found", func(t *testing.T) {
		_, err := GetLongURL("123456")
		assert.EqualError(t, domain.ErrURLNotFound, err.Error())
	})
}
