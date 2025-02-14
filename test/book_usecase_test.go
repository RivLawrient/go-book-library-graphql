package test

import (
	"fmt"
	"go-book-library-graphql/internal/app/book"
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
	result, err := book.NewBook(&req)

	assert.Nil(t, err)
	assert.Equal(t, req.Title, result.Title)
	assert.Equal(t, req.Author, result.Author)
	assert.Equal(t, req.Language, result.Language)
	assert.Equal(t, *req.Synopsis, result.Synopsis)
	assert.Equal(t, req.TotalPage, result.TotalPage)
}

func TestShowALlBook(t *testing.T) {
	result, err := book.ShowAllBook()
	assert.Nil(t, err)
	fmt.Println(*result)

}
