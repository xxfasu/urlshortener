package response

import "github.com/xxfasu/urlshortener/internal/common/data"

type CreateURL struct {
	ShortURL string `json:"short_url"`
}

type GetURLs struct {
	Items []data.FullURL `json:"items"`
	Total int            `json:"total"`
}
