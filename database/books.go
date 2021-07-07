package database

import (
	"database/sql"

	"github.com/gussf/go-bookstore/model"
)

func (c *Connection) SelectAllBooks() ([]model.Book, error) {

	stmt, err := c.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	var bookList []model.Book
	for stmt.Next() {
		book, err := scanBookFrom(stmt)
		if err != nil {
			return nil, err
		}
		bookList = append(bookList, *book)
	}

	return bookList, nil
}

func (c *Connection) SelectBookById(id string) (*model.Book, error) {

	stmt, err := c.db.Query("SELECT * FROM books WHERE id = " + id)
	if err != nil {
		return nil, err
	}

	if !stmt.Next() {
		return nil, nil
	}

	book, err := scanBookFrom(stmt)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func scanBookFrom(stmt *sql.Rows) (*model.Book, error) {
	var book model.Book
	err := stmt.Scan(&book.ID, &book.Title, &book.Author, &book.Copies, &book.Price, &book.CreationDate)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
