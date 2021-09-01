package repository

import (
	"database/sql"
	"fmt"
	"time"

	bookstore "github.com/gussf/go-bookstore/src"
)

// Postgres repository
type pgBookRepository struct {
	Conn *PostgresConnection
}

func NewPostgresRepo() (*pgBookRepository, error) {
	c, err := NewPostgresConnection()
	if err != nil {
		return nil, err
	}
	return &pgBookRepository{Conn: c}, nil
}

func (br pgBookRepository) NewBook(title string, author string, copies int, price int64) *bookstore.BookDTO {
	return &bookstore.BookDTO{
		Title:  title,
		Author: author,
		Copies: copies,
		Price:  price,
	}
}

func (br pgBookRepository) SelectAll() ([]bookstore.BookDTO, error) {

	stmt, err := br.Conn.DB.Query("SELECT * FROM books LIMIT 100")
	if err != nil {
		return nil, err
	}

	var bookList []bookstore.BookDTO
	for stmt.Next() {
		book, err := scanBookFrom(stmt)
		if err != nil {
			return nil, err
		}
		bookList = append(bookList, *book)
	}

	return bookList, nil
}

func (br pgBookRepository) Select(id string) (*bookstore.BookDTO, error) {

	stmt, err := br.Conn.DB.Query("SELECT * FROM books WHERE id = " + id)
	if err != nil {
		return nil, err
	}

	// Book not found
	if !stmt.Next() {
		return nil, fmt.Errorf("{}")
	}

	book, err := scanBookFrom(stmt)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (br pgBookRepository) Insert(b *bookstore.BookDTO) error {
	lastInsertedId := 0
	var creationDate time.Time
	row := br.Conn.DB.QueryRow("insert into books(title, author, copies, price, creation_date) values($1,$2,$3,$4, current_timestamp) RETURNING id, creation_date", b.Title, b.Author, b.Copies, b.Price)
	err := row.Scan(&lastInsertedId, &creationDate)
	if err != nil {
		return err
	}
	b.ID = lastInsertedId
	b.CreationDate = creationDate
	return nil
}

func (br pgBookRepository) Delete(id string) error {
	res, err := br.Conn.DB.Exec("delete from books where id=$1", id)
	if err != nil {
		return err
	}

	r, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("DELETE - Rows affected: ", r)
	return nil
}

func scanBookFrom(stmt *sql.Rows) (*bookstore.BookDTO, error) {
	var book bookstore.BookDTO
	err := stmt.Scan(&book.ID, &book.Title, &book.Author, &book.Copies, &book.Price, &book.CreationDate)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (br pgBookRepository) CloseConnection() {
	_ = br.Conn.Close()
}
