package resolver

import (
	"go-book-library-graphql/internal/app/book"

	"github.com/graphql-go/graphql"
)

type BookResolver struct {
	BookUsecase *book.BookUsecase
}

func NewBookResolver(bookUsecase *book.BookUsecase) *BookResolver {
	return &BookResolver{
		BookUsecase: bookUsecase,
	}
}

func (b *BookResolver) NewBookHandler(param graphql.ResolveParams) (interface{}, error) {
	return "book newBookHandler", nil
}
