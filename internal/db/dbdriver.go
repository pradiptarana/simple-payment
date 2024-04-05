package db

import (
	"database/sql"
	"os"

	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewDBConnection() *sql.DB {
	// Capture connection properties.
	// cfg := mysql.Config{
	// 	User:   os.Getenv("DB_USER"),
	// 	Passwd: os.Getenv("DB_PASS"),
	// 	Net:    "tcp",
	// 	Addr:   os.Getenv("DB_ADDR"),
	// 	DBName: os.Getenv("DB_NAME"),
	// }
	cfg := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_ADDR"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println(cfg)

	// Get a database handle.
	var err error
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}
