package domain

type UrlRecord struct {
	ID      string `json:"id"`
	LongUrl string `json:"longUrl"`
}

type UrlInfo struct {
	Url string
}
