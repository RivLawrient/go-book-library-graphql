package book

import (
	"database/sql"
	"go-book-library-graphql/internal/config"
	"log"
)

func Create(id string, req *NewBookRequest) error {
	log.Println("BookCreateRepo")
	query := "INSERT INTO book(id, title, author, language, synopsis, total_page) VALUES($1, $2, $3, $4, $5, $6)"
	_, err := config.GetConnection().Exec(query, &id, &req.Title, &req.Author, &req.Language, &req.Synopsis, &req.TotalPage)

	return err
}

func FindAll() (*[]Book, error) {
	query := "SELECT id, title, author, language, synopsis, total_page, current_page, created_at FROM book"
	result, err := config.GetConnection().Query(query)
	if err != nil {
		return nil, err
	}
	datas := []Book{}

	for result.Next() {
		data := Book{}
		err := result.Scan(&data.Id, &data.Title, &data.Author, &data.Language, &data.Synopsis, &data.TotalPage, &data.CurrentPage, &data.CreatedAt)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}

	if len(datas) == 0 {
		return nil, sql.ErrNoRows
	}
	return &datas, nil
}
