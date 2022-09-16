package models

type URL struct {
	ShortURL    string `json:"short_url" db:"short_url"`
	OriginalURL string `json:"original_url" db:"full_url"`
}

type URLs []URL
