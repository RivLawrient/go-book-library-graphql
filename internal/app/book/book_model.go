package book

import "time"

type NewBookResponse struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Language  string    `json:"language"`
	Synopsis  string    `json:"synopsis"`
	TotalPage int       `json:"total_page"`
	CreatedAt time.Time `json:"created_at"`
}

type NewBookRequest struct {
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Language  string  `json:"language"`
	Synopsis  *string `json:"synopsis"`
	TotalPage int     `json:"total_page"`
}
