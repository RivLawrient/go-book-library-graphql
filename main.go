package main

import (
	"fmt"
	"go-book-library-graphql/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.Handle("/graphql", handlers.GraphQLHandler())

	fmt.Println("🚀 Server berjalan di http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
