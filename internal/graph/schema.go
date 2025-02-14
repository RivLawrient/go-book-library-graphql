package graph

import (
	"go-book-library-graphql/internal/graph/resolver"

	"github.com/graphql-go/graphql"
)

type SchemaConfig struct {
	BookResolver *resolver.BookResolver
}

func (s *SchemaConfig) Route() *graphql.Schema {
	// Root Query
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Type:    graphql.String,
				Resolve: s.BookResolver.NewBookHandler,
			},
		},
	})

	// Root Mutation
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"create": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					res := p.Args["name"]
					return res, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	return &schema

}
