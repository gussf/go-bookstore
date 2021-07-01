package model_test

import (
	"testing"

	"github.com/gussf/bookstore/internal/model"
)

func TestBook(t *testing.T) {
	_ = model.Book{
		Title:  "Spark Joy",
		Author: []string{"Marie Kondō"},
		Copies: 3,
		Price:  15,
	}
}
