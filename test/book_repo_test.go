package test

import (
	"go-book-library-graphql/internal/app/book"
	"go-book-library-graphql/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	synopsis := "kosong"
	req := book.NewBookRequest{
		Title:     "rumah",
		Author:    "gatau",
		Language:  "eng",
		Synopsis:  &synopsis,
		TotalPage: 12,
	}
	result, err := book.NewBookUsecase(book.NewBookRepository(config.GetConnection())).NewBook(&req)

	assert.Nil(t, err)
	assert.Equal(t, req.Title, result.Title)
	assert.Equal(t, req.Author, result.Author)
	assert.Equal(t, req.Language, result.Language)
	assert.Equal(t, *req.Synopsis, result.Synopsis)
	assert.Equal(t, req.TotalPage, result.TotalPage)
}
