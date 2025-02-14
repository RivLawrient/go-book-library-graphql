package book

import "database/sql"

type BookRepository struct {
	Db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		Db: db,
	}
}
