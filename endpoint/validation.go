package endpoint

import (
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"url-shortener/domain"

	"github.com/google/uuid"
)

func IsURLValid(url string) error {
	_, err := http.Get(url)
	if err != nil {
		return err
	}
	return nil
}

func IsURLAleadyShortened(host string, requestUrl string) error {
	parsedUrl, err := url.Parse(requestUrl)
	if err != nil {
		return err
	}
	if strings.Contains(parsedUrl.String(), fmt.Sprintf("%s/%s", host, domain.ShortUrl)) {
		// Find urlRecordID
		_, urlRecordID := filepath.Split(parsedUrl.Path)
		if IsUUIDValid(urlRecordID) {
			return domain.ErrURLIsAlreadyShorted
		}
	}
	return nil
}

func IsUUIDValid(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
