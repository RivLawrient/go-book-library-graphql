package resolver

import "go-book-library-graphql/internal/app/book"

type NewBookResolver struct {
	BookUsecase *book.BookUsecase
}


