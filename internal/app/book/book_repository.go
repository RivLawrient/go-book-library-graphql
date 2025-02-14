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

func (b *BookRepository) Create(id string, req *NewBookRequest) error {
	query := "INSERT INTO book(id, title, author, language, synopsis, total_page) VALUES($1, $2, $3, $4, $5, $6, $7)"
	_, err := b.Db.Exec(query, id, req.Title, req.Author, req.Language, req.Synopsis, req.TotalPage)

	return err
}
