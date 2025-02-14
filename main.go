package main

import (
	"fmt"
	"go-book-library-graphql/internal/config"
	"go-book-library-graphql/internal/handlers"
	"log"
	"net/http"
)

func main() {
	db := config.GetConnection()
	reg := &config.RegisterConfig{
		Db: db,
	}
	http.Handle("/graphql", handlers.GraphQLHandler(config.Register(reg)))

	fmt.Println("🚀 Server berjalan di http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
