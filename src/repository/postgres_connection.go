package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type PostgresConnection struct {
	DB *sql.DB
}

// Connects to a database according to the following environment variables: DRIVER, USER, HOST, PORT, PASSWORD, DBNAME.
// Example:  DRIVER=postgres, USER=db_user, HOST=localhost, PORT=5432, PASSWORD=12345, DBNAME=database_name
func NewPostgresConnection() (c *PostgresConnection, err error) {

	fmt.Println("Connecting to PostgreSQL database ...")

	driver := os.Getenv("DRIVER")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println(connectionString)
	db, err := sql.Open(driver, connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected!")

	return &PostgresConnection{db}, nil
}

// Closes a previously instantiated database connection
func (c *PostgresConnection) Close() error {
	return c.DB.Close()
}
