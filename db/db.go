package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST = "db"
	PORT = 5432
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

// Initialize database...
func Initialize(username, password, database string) (Database, error) {
	//retries := 0
	//timeout := 5

	//for retries <= 5 {
	log.Println("Database connection initiating...")
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("Connection Error")
		log.Println(err)
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		log.Println("Ping error")
		log.Println(err)
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
	//}

}
