package models

import "time"

type ClickEvent struct {
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
	Referrer  string    `json:"referrer"`
}

type Link struct {
	ID          string       `json:"id"`
	Slug        string       `json:"slug"`
	OriginalURL string       `json:"original_url"`
	Metadata    []byte       `json:"metadata"`
	Analytics   []ClickEvent `json:"analytics"`
	CreatedAt   time.Time    `json:"created_at"`
}
