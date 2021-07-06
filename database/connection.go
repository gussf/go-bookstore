package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Connection struct {
	db *sql.DB
}

// NewConnection
// Connects to a database according to the following environment variables:
// DRIVER  		ex: postgres
// USER    		ex: db_user
// HOST   	 	ex: localhost
// PORT	   		ex: 5432
// PASSWORD		ex: 12345
// DBNAME 		ex: database_name
func NewConnection() (c *Connection, err error) {

	fmt.Println("Connecting to database ...")

	driver := os.Getenv("DRIVER")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open(driver, connectionString)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected!")

	return &Connection{db}, nil
}

// Close
// Closes a previously instantiated database connection
func (c *Connection) Close() error {
	return c.db.Close()
}
