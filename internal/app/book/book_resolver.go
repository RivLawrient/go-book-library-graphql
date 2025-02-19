package book

import (
	"github.com/graphql-go/graphql"
)

var NewBookMutation = &graphql.Field{
	Type: NewBookResponseModel,
	Args: graphql.FieldConfigArgument{
		"title":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"author":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"language":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"synopsis":   &graphql.ArgumentConfig{Type: graphql.String},
		"total_page": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		synopsis, found := p.Args["synopsis"].(string)
		if !found {
			synopsis = ""
		}
		req := &NewBookRequest{
			Title:     p.Args["title"].(string),
			Author:    p.Args["author"].(string),
			Language:  p.Args["language"].(string),
			Synopsis:  &synopsis,
			TotalPage: p.Args["total_page"].(int),
		}

		result, err := NewBook(req)
		if err != nil {
			return nil, err
		}

		return result, nil
	},
}

var ShowAllBookQuery = &graphql.Field{
	Type: graphql.NewList(BookModel),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		result, err := ShowAllBook()
		if err != nil {
			return nil, err
		}

		return result, nil
	},
}

var ShowByIdBookQuery = &graphql.Field{
	Type: BookModel,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(string)
		result, err := ShowById(id)
		if err != nil {
			return nil, err
		}

		return result, nil
	},
}

var RemoveByIdBookMutation = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(string)
		err := RemoveBookById(id)
		if err != nil {
			return nil, err
		}

		return "Success remove " + id, nil
	},
}
