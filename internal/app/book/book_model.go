package book

import (
	"time"

	"github.com/graphql-go/graphql"
)

type Book struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Language    string    `json:"language"`
	Synopsis    string    `json:"synopsis"`
	TotalPage   int       `json:"total_page"`
	CurrentPage int       `json:"current_page"`
	CreatedAt   time.Time `json:"created_at"`
}

type NewBookResponse struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Language  string    `json:"language"`
	Synopsis  string    `json:"synopsis"`
	TotalPage int       `json:"total_page"`
	CreatedAt time.Time `json:"created_at"`
}

type NewBookRequest struct {
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Language  string  `json:"language"`
	Synopsis  *string `json:"synopsis"`
	TotalPage int     `json:"total_page"`
}

var BookModel = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id":           &graphql.Field{Type: graphql.String},
		"title":        &graphql.Field{Type: graphql.String},
		"author":       &graphql.Field{Type: graphql.String},
		"language":     &graphql.Field{Type: graphql.String},
		"synopsis":     &graphql.Field{Type: graphql.String},
		"total_page":   &graphql.Field{Type: graphql.Int},
		"current_page": &graphql.Field{Type: graphql.Int},
		"created_at":   &graphql.Field{Type: graphql.DateTime},
	},
})

var NewBookResponseModel = graphql.NewObject(graphql.ObjectConfig{
	Name: "NewBookResponse",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.String},
		"title":      &graphql.Field{Type: graphql.String},
		"author":     &graphql.Field{Type: graphql.String},
		"language":   &graphql.Field{Type: graphql.String},
		"synopsis":   &graphql.Field{Type: graphql.String},
		"total_page": &graphql.Field{Type: graphql.Int},
		"created_at": &graphql.Field{Type: graphql.DateTime},
	},
})
