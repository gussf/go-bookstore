package bookstore_test

import (
	"reflect"
	"testing"

	"github.com/gussf/bookstore"
)

func init() {
	bookstore.Catalog = []bookstore.Book{testB}
}

var testB = bookstore.Book{
	Title:       "Spark Joy",
	Author:      []string{"Marie Kondō"},
	Description: "A tiny, cheerful Japanese woman explains tidying.",
	PriceCents:  1199,
	Featured:    true,
}

func TestBook(t *testing.T) {
	_ = bookstore.Book{
		Title:       "Spark Joy",
		Author:      []string{"Marie Kondō"},
		Description: "A tiny, cheerful Japanese woman explains tidying.",
		PriceCents:  1199,
	}
}

func TestAddToCatalog(t *testing.T) {

	catalog := bookstore.AddToCatalog(testB)
	if !reflect.DeepEqual(catalog[0], testB) {
		t.Errorf("AddToCatalog - want %v, got %v", testB, catalog[0])
	}
}

func TestFeaturedBooks(t *testing.T) {
	books := bookstore.FeaturedBooks()

	if len(books) == 0 {
		t.Errorf("FeaturedBooks - want 1, got 0")
	}
}
