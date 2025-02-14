package book

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func NewBook(req *NewBookRequest) (*NewBookResponse, error) {
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
