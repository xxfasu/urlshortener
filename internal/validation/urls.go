package validation

import "time"

type CreateURL struct {
	OriginalURL string `json:"original_url" binding:"required,url"`
	CustomCode  string `json:"custom_code,omitempty" binding:"omitempty,min=4,max=10,alphanum"`
	Duration    *int   `json:"duration,omitempty" binding:"omitempty,min=1,max=100"`
	UserID      int    `json:"-"`
}

type GetURLs struct {
	Page   uint `query:"page"`
	Size   uint `query:"size"`
	UserID int  `query:"-"`
}

type DeleteURL struct {
	Code string `param:"code" binding:"required,len=6,alphanum"`
}

type UpdateURLDuration struct {
	Code      string    `param:"code" binding:"required,len=6,alphanum"`
	ExpiredAt time.Time `json:"expired_at" binding:"required,after"`
}
