package service

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/PrateekSawala/url-shortener/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateShortURLSuccess(t *testing.T) {

	t.Run("Should return a new generated short Url", func(t *testing.T) {
		response, err := GenerateShortURL("localhost:3002", fmt.Sprintf("https://%d.com", time.Now().UnixNano()))
		if err != nil {
			t.Errorf("GenerateShortURL Error: %s", err)
			return
		}
		// Find urlRecordID
		_, urlRecordID := filepath.Split(response.Url)
		_, err = uuid.Parse(urlRecordID)
		assert.Equal(t, err == nil, true)
	})

	t.Run("Should return a already shortened URL", func(t *testing.T) {
		response, err := GenerateShortURL("localhost:3002", "https://google.com")
		if err != nil {
			t.Errorf("GenerateShortURL Error: %s", err)
			return
		}
		result := &domain.UrlInfo{Url: "http://localhost:3002/shortUrl/2aa789c823074703b7baa8a524eb4aad"}
		assert.Equal(t, result, response)
	})
}

func TestShortURLGeneratorError(t *testing.T) {
	_, err := GenerateShortURL("", "")
	assert.EqualError(t, domain.ErrInvalidInput, err.Error())
}
