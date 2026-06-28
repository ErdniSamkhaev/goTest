package main

import (
	"time"
)

// структура для бд
type Link struct {
	ID          int       `json:"id"`
	ShortCode   string    `json:"short_code"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}

// входящий запрос
type ShortenRequest struct {
	URL string `json:"url"`
}

// Исходящий запрос
type ShortenResponse struct {
	ShortCode string `json:"short_code"`
	ShortUrl  string `json:"short_url"`
}
