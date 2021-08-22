package database

import (
	"database/sql"
	"fmt"

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

func (c *Connection) InsertBook(b *model.Book) error {
	lastInsertedId := 0
	row := c.db.QueryRow("insert into books(title, author, copies, price, creation_date) values($1,$2,$3,$4, current_timestamp) RETURNING id", b.Title, b.Author, b.Copies, b.Price)
	err := row.Scan(&lastInsertedId)
	if err != nil {
		return err
	}
	fmt.Printf("New book inserted with id=%d\n", lastInsertedId)
	b.ID = lastInsertedId
	return nil
}
