package data

import "time"

type URL struct {
	OriginalURL string
	ShortCode   string
}

type FullURL struct {
	ID          int       `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	ExpiredAt   time.Time `json:"expired_at"`
	IsCustom    bool      `json:"is_custom"`
	Views       uint      `json:"views"`
}
