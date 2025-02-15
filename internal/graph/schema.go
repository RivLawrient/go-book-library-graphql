package graph

import (
	"go-book-library-graphql/internal/app/book"

	"github.com/graphql-go/graphql"
)

func Schema() *graphql.Schema {
	// Root Query
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"book": book.ShowAllBookQuery,
		},
	})

	// Root Mutation
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"newBook": book.NewBookMutation,
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	return &schema

}
