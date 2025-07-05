package Config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func (c connection) PostgreSQLConnection() *sql.DB {
	return c.postgresql
}

// build connection
func newConnectionPostgreSQL() *sql.DB {
	host := os.Getenv("HOST_DB")
	if host == "" {
		host = "localhost"
	}

	db, err := sql.Open("postgres", "postgres://postgres:root1234@"+host+":5432/article-kump?sslmode=disable")
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	fmt.Println("Connected to PostgreSQL")

	return db
}
