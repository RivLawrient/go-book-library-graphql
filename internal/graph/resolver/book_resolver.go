package resolver

import (
	"github.com/graphql-go/graphql"
)

// type BookResolver struct {
// 	BookUsecase *book.BookUsecase
// }

// func NewBookResolver(bookUsecase *book.BookUsecase) *BookResolver {
// 	return &BookResolver{
// 		BookUsecase: bookUsecase,
// 	}
// }

func NewBookHandler(param graphql.ResolveParams) (interface{}, error) {
	return "book newBookHandler", nil
}
