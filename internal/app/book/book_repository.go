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
	log.Println("FindAllRepo")
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

func FindById(id string) (*Book, error) {
	log.Println("FindByIdRepo")
	query := "SELECT title, author, language, synopsis, total_page, current_page, created_at FROM book WHERE id=$1"
	result := config.GetConnection().QueryRow(query, id)
	if result.Err() != nil {
		return nil, result.Err()
	}

	data := new(Book)
	data.Id = id
	if err := result.Scan(&data.Title, &data.Author, &data.Language, &data.Synopsis, &data.TotalPage, &data.CurrentPage, &data.CreatedAt); err != nil {
		return nil, err
	}

	return data, nil
}

func RemoveById(id string) error {
	query := "DELETE FROM book WHERE id=$1;"
	result, err := config.GetConnection().Exec(query, id)
	if err != nil {
		return err
	}

	if n, _ := result.RowsAffected(); n == 0 {
		return sql.ErrNoRows
	}
	return err
}
