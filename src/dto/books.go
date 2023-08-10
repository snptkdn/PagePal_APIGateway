package dto

import "time"

type BookRequestDto struct {
	ISBN        string `json:"ISBN"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	PageCount   int    `json:"page_count"`
}

type ReadHistoryRequestDto struct {
	BookID uint      `json:"book_id"`
	UserID uint      `json:"user_id"`
	IsRead bool      `json:"is_read"`
	Rate   int       `json:"rate"`
	Date   time.Time `json:"date"`
}
