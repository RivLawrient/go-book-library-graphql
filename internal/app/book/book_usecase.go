package book

import (
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

func NewBook(req *NewBookRequest) (*NewBookResponse, error) {
	log.Println("NewBookUsecase")
	id := uuid.New().String()
	req.Title = strings.TrimSpace(req.Title)
	err := Create(id, req)
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

func ShowAllBook() (*[]Book, error) {
	log.Println("ShowAllBookUsecase")
	result, err := FindAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ShowById(id string) (*Book, error) {
	log.Println("ShowByIdUsecase")
	result, err := FindById(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
