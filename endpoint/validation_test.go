package endpoint

import (
	"fmt"
	"testing"

	"url-shortener/domain"

	"github.com/stretchr/testify/assert"
)

func TestIsURLValidSuccess(t *testing.T) {
	// Declare url
	url := fmt.Sprintf("https://google.com")
	err := IsURLValid(url)
	assert.NoError(t, err)
}

func TestIsURLAleadyShortenedError(t *testing.T) {
	// Declare url
	url := fmt.Sprintf("http://localhost:3002/shortUrl/2aa789c823074703b7baa8a524eb4aad")
	err := IsURLAleadyShortened("http://localhost:3002", url)
	expectedErr := domain.ErrURLIsAlreadyShorted
	assert.EqualError(t, expectedErr, err.Error())
}

func TestIsUUIDValidSuccess(t *testing.T) {
	valid := IsUUIDValid("2aa789c823074703b7baa8a524eb4aad")
	assert.Equal(t, true, valid)
}
