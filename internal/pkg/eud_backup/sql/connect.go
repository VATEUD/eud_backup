package sql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(name string) (*sql.DB, error) {

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, name))

	if err != nil {
		log.Println(err)
		return &sql.DB{}, err
	}

	log.Println(db)

	return db, nil
}
