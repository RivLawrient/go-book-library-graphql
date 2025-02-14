package handlers

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func GraphQLHandler(schema *graphql.Schema) http.Handler {
	h := handler.New(&handler.Config{
		Schema:     schema,
		Pretty:     true,
		GraphiQL:   true, // Aktifkan GraphiQL di browser
		Playground: true, // Aktifkan Playground
	})
	return h
}
