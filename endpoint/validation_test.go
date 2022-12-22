package endpoint

import (
	"fmt"
	"testing"

	"github.com/PrateekSawala/url-shortener/domain"

	"github.com/stretchr/testify/assert"
)

func TestIsURLValidSuccess(t *testing.T) {
	t.Run("Should return no error", func(t *testing.T) {
		err := IsURLValid("http://localhost:3002", "https://google.com")
		assert.NoError(t, err)
	})
}

func TestIsURLValidError(t *testing.T) {
	t.Run("Should return error of url is already shorted", func(t *testing.T) {
		// Declare url
		url := fmt.Sprintf("http://localhost:3002/shortUrl/2aa789c823074703b7baa8a524eb4aad")
		err := IsURLValid("http://localhost:3002", url)
		expectedErr := domain.ErrURLIsAlreadyShortened
		assert.EqualError(t, expectedErr, err.Error())
	})
	t.Run("Should return error of invalid url", func(t *testing.T) {
		err := IsURLValid("http://localhost:3002", "")
		expectedErr := domain.ErrInvalidURL
		assert.EqualError(t, expectedErr, err.Error())
	})
	t.Run("Should return error of invalid url", func(t *testing.T) {
		err := IsURLValid("http://localhost:3002", "https://")
		expectedErr := domain.ErrInvalidURL
		assert.EqualError(t, expectedErr, err.Error())
	})
}

func TestIsUUIDValid(t *testing.T) {
	t.Run("Should return true", func(t *testing.T) {
		valid := IsUUIDValid("2aa789c823074703b7baa8a524eb4aad")
		assert.Equal(t, valid, true)
	})
	t.Run("Should return false", func(t *testing.T) {
		valid := IsUUIDValid("2aa789c82")
		assert.Equal(t, valid, false)
	})
}
