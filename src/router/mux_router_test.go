package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	bookstore "github.com/gussf/go-bookstore/src"
	"github.com/gussf/go-bookstore/src/books"
	"github.com/gussf/go-bookstore/src/repository"
)

var repo bookstore.Repository
var svc books.Service
var r bookstore.Router

var Books map[string]books.Book
var realLength int

func TestMain(m *testing.M) {
	setup()

	code := bookstore.Run()
	os.Exit(code)
}

func setup() {

	Books = map[string]bookstore.BookDTO{
		"1": {ID: 1, Title: "Game of Thrones", Author: "George", Copies: 1, Price: 3000, CreationDate: time.Now()},
		"2": {ID: 2, Title: "Alice in Wonderland", Author: "Lewis", Copies: 5, Price: 1500, CreationDate: time.Now()},
	}
	realLength = len(Books)

	repo = repository.InMemRepository{
		BookList: Books,
	}

	svc = books.NewService(repo)
	r = NewMuxRouter(svc)

}

func TestMuxRouter_FindBookByID(t *testing.T) {

	tests := []struct {
		name               string
		id                 int
		expectedHTTPStatus int
		expectedTitle      string
	}{
		{
			name:               "Find existing Book by id",
			id:                 1,
			expectedHTTPStatus: http.StatusOK,
			expectedTitle:      "Game of Thrones",
		},
		{
			name:               "Find nonexistent book by id",
			id:                 40,
			expectedHTTPStatus: http.StatusNotFound,
			expectedTitle:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/book/%d", tt.id), nil)
			response := httptest.NewRecorder()

			r.ServeHTTP(response, request)

			var found bookstore.Book
			_ = json.Unmarshal(response.Body.Bytes(), &found)

			assertAttribute(t, "Title", found.Title, tt.expectedTitle)
			assertStatus(t, response.Code, tt.expectedHTTPStatus)

		})
	}
}

func TestMuxRouter_FindAllBooks(t *testing.T) {

	tests := []struct {
		name               string
		expectedHTTPStatus int
		expectedLength     int
	}{
		{
			name:               "Find all books",
			expectedLength:     realLength,
			expectedHTTPStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/api/v1/books", nil)
			response := httptest.NewRecorder()

			r.ServeHTTP(response, request)

			var found []bookstore.Book
			_ = json.Unmarshal(response.Body.Bytes(), &found)

			if len(found) != tt.expectedLength {
				t.Errorf("did not get correct number of books, got %d, want %d", len(found), tt.expectedLength)
			}

			assertStatus(t, response.Code, tt.expectedHTTPStatus)

		})
	}
}

func TestMuxRouter_DeleteBookById(t *testing.T) {

	tests := []struct {
		name                  string
		expectedContentLength string
		expectedHTTPStatus    int
		id                    int
	}{
		{
			name:                  "Delete existing book with id=2",
			expectedContentLength: "",
			expectedHTTPStatus:    http.StatusNoContent,
			id:                    2,
		},
		{
			name:                  "Delete nonexistent book with id=99",
			expectedContentLength: "",
			expectedHTTPStatus:    http.StatusNoContent,
			id:                    99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/book/%d", tt.id), nil)
			response := httptest.NewRecorder()

			r.ServeHTTP(response, request)

			assertAttribute(t, "Content-Length", response.Header().Get("Content-Length"), tt.expectedContentLength)
			assertStatus(t, response.Code, tt.expectedHTTPStatus)
		})
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct http status, got %d, want %d", got, want)
	}
}

func assertAttribute(t testing.TB, attribute, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("["+attribute+"] is wrong, got %q want %q", got, want)
	}
}
