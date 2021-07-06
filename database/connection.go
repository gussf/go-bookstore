package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gussf/go-bookstore/model"

	_ "github.com/lib/pq"
)

type Connection struct {
	db *sql.DB
}

// Connects to a database according to the following environment variables: DRIVER, USER, HOST, PORT, PASSWORD, DBNAME.
// Example:  DRIVER=postgres, USER=db_user, HOST=localhost, PORT=5432, PASSWORD=12345, DBNAME=database_name
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

// Closes a previously instantiated database connection
func (c *Connection) Close() error {
	return c.db.Close()
}

func (c *Connection) SelectAllBooks() ([]model.Book, error) {

	stmt, err := c.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	var bookList []model.Book
	for stmt.Next() {
		var book model.Book
		err = stmt.Scan(&book.ID, &book.Title, &book.Author, &book.Copies, &book.Price, &book.CreationDate)
		if err != nil {
			return nil, err
		}
		bookList = append(bookList, book)
	}

	return bookList, nil
}
