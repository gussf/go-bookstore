package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gussf/go-bookstore/database"
	"github.com/gussf/go-bookstore/model"
)

// Postgres repository
type PgBookRepository struct {
	Conn *database.PostgresConnection
}

func NewPostgresRepo() (*PgBookRepository, error) {
	c, err := database.NewPostgresConnection()
	if err != nil {
		return nil, err
	}
	return &PgBookRepository{Conn: c}, nil
}

func (br PgBookRepository) SelectAll() ([]model.Book, error) {

	stmt, err := br.Conn.DB.Query("SELECT * FROM books LIMIT 100")
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

func (br PgBookRepository) Select(id string) (*model.Book, error) {

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

func (br PgBookRepository) Insert(b *model.Book) error {
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

func (br PgBookRepository) Delete(id string) error {
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

func scanBookFrom(stmt *sql.Rows) (*model.Book, error) {
	var book model.Book
	err := stmt.Scan(&book.ID, &book.Title, &book.Author, &book.Copies, &book.Price, &book.CreationDate)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
