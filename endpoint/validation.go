package endpoint

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/PrateekSawala/url-shortener/domain"

	"github.com/google/uuid"
)

func IsURLValid(host string, requestUrl string) error {
	if host == "" || requestUrl == "" {
		return domain.ErrInvalidURL
	}
	parsedUrl, err := url.Parse(requestUrl)
	if err != nil {
		return err
	}
	if IsURLAlreadyShorted(host, parsedUrl) {
		return domain.ErrURLIsAlreadyShortened
	}
	if err == nil && parsedUrl.Scheme != "" && parsedUrl.Host != "" {
		return nil
	}
	return domain.ErrInvalidURL
}

func IsURLAlreadyShorted(host string, parsedUrl *url.URL) bool {
	if strings.Contains(parsedUrl.String(), fmt.Sprintf("%s/%s", host, domain.ShortUrl)) {
		// Find urlRecordID
		_, urlRecordID := filepath.Split(parsedUrl.Path)
		if IsUUIDValid(urlRecordID) {
			return true
		}
	}
	return false
}

func IsUUIDValid(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
