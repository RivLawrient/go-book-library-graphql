package config

import (
	"database/sql"
	"go-book-library-graphql/internal/app/book"
	"go-book-library-graphql/internal/graph"
	"go-book-library-graphql/internal/graph/resolver"

	"github.com/graphql-go/graphql"
)

type RegisterConfig struct {
	Db *sql.DB
}

func Register(cfg *RegisterConfig) *graphql.Schema {
	bookRepository := book.NewBookRepository(cfg.Db)
	bookUsecase := book.NewBookUsecase(bookRepository)
	bookResolver := resolver.NewBookResolver(bookUsecase)

	config := graph.SchemaConfig{
		BookResolver: bookResolver,
	}

	return config.Route()
}
