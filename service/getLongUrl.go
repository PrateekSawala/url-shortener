package service

import (
	"url-shortener/domain/logging"

	"url-shortener/domain"
)

func GetLongURL(shorturl string) (*domain.UrlInfo, error) {
	log := logging.Log("GetLongURL")
	log.Tracef("Start")
	defer log.Tracef("End")

	response := &domain.UrlInfo{}
	URLRecords, err := GetURLRecords()
	if err != nil {
		log.Debugf("GetURLRecords error: %s", err)
		return response, err
	}
	urlRecord := CheckUrl(URLRecords, shorturl)
	// Check urlRecord
	if urlRecord.LongUrl == "" {
		log.Debugf("%s", domain.ErrURLNotFound)
		return response, domain.ErrURLNotFound
	}
	response.Url = urlRecord.LongUrl
	return response, nil
}