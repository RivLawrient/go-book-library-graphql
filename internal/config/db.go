package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	// username := viper.GetString("DB_USERNAME")
	// password := viper.GetString("DB_PASSWORD")
	// host := viper.GetString("DB_HOST")
	// port := viper.GetInt("DB_PORT")
	// database := viper.GetString("DB_NAME")
	// maxIdleConnection := viper.GetInt("DB_MAX_IDLE_CONSS")
	// maxOpenConns := viper.GetInt("DB_MAX_OPEN_CONNS")
	// connMaxIdleTime := viper.GetInt("DB_CONN_MAX_IDLE") // in minute
	// connMaxLifetime := viper.GetInt("DB_CONN_MAX_LIFE") // in minute

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	maxIdleConnection, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	maxOpenConns, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	connMaxIdleTime, _ := strconv.Atoi(os.Getenv("DB_CONN_MAX_IDLE"))
	connMaxLifetime, _ := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFE"))

	var dsn string
	if password == "" {
		// Jika password kosong, tidak sertakan parameter password
		dsn = fmt.Sprintf(
			"host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Makassar",
			host, username, database, port, "disable",
		)
	} else {
		// Jika password ada, sertakan parameter password
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Makassar",
			host, username, password, database, port, "disable",
		)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(maxIdleConnection)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxIdleTime(time.Duration(connMaxIdleTime) * time.Minute)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Minute)

	return db
}
