package book

import "time"

type Book struct {
	Id          string
	Title       string
	Author      string
	Language    string
	Synopsis    string
	TotalPage   int
	CurrentPage int
	CreatedAt   time.Time
}
