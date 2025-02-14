package book

import (
	"go-book-library-graphql/internal/config"
)

func Create(id string, req *NewBookRequest) error {
	query := "INSERT INTO book(id, title, author, language, synopsis, total_page) VALUES($1, $2, $3, $4, $5, $6)"
	_, err := config.GetConnection().Exec(query, id, req.Title, req.Author, req.Language, req.Synopsis, req.TotalPage)

	return err
}
