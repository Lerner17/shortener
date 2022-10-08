package models

type URL struct {
	ShortURL    string `json:"short_url" db:"short_url"`
	OriginalURL string `json:"original_url" db:"full_url"`
}

type URLs []URL

type BatchURL struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type BatchURLs []BatchURL

type BatchShortURL struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url" db:"short_url"`
}

type BatchShortURLs []BatchShortURL

type URLEntity struct {
	OriginURL     string
	ShortURL      string
	UserSession   string
	CorrelationID string
	IsDeleted     bool
}
