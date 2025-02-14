package test

import (
	"go-book-library-graphql/internal/config"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	db := config.GetConnection()

	assert.Nil(t, db.Ping())

}
