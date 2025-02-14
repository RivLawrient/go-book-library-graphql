package handlers

import (
	"go-book-library-graphql/internal/graph"
	"net/http"

	"github.com/graphql-go/handler"
)

func GraphQLHandler() http.Handler {
	h := handler.New(&handler.Config{
		Schema:     &graph.Schema,
		Pretty:     true,
		GraphiQL:   true, // Aktifkan GraphiQL di browser
		Playground: true, // Aktifkan Playground
	})
	return h
}
