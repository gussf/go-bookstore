package model_test

import (
	"testing"

	"github.com/gussf/go-bookstore/model"
)

func TestBook(t *testing.T) {
	_ = model.Book{
		Title:  "Spark Joy",
		Author: []string{"Marie Kondō"},
		Copies: 3,
		Price:  15,
	}
}
