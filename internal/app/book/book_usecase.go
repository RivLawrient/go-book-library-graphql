package book

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type BookUsecase struct {
	BookRepository *BookRepository
}

func NewBookUsecase(bookRepo *BookRepository) *BookUsecase {
	return &BookUsecase{
		BookRepository: bookRepo,
	}
}

func (b *BookUsecase) NewBook(req *NewBookRequest) (*NewBookResponse, error) {
	id := uuid.New().String()
	req.Title = strings.TrimSpace(req.Title)
	err := b.BookRepository.Create(id, req)
	if err != nil {
		return nil, err
	}

	return &NewBookResponse{
		Id:        id,
		Title:     req.Title,
		Author:    req.Author,
		Language:  req.Language,
		Synopsis:  *req.Synopsis,
		TotalPage: req.TotalPage,
		CreatedAt: time.Now(),
	}, nil
}
