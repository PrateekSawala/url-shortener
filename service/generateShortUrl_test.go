package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortURLGeneratorSuccess(t *testing.T) {
	_, err := GenerateShortURL("localhost:3002", "https://google.com")
	if err != nil {
		t.Errorf("GenerateShortURL Error: %s", err)
		return
	}
	assert.NoError(t, err)
}
