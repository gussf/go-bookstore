package books

import (
	"reflect"
	"testing"
)

func TestNewBook(t *testing.T) {
	type args struct {
		title  string
		author string
		copies int
		price  int64
	}
	tests := []struct {
		name     string
		args     args
		wantBook *Book
		wantErr  error
	}{
		{
			name: "Valid book",
			args: args{
				title:  "Good Title",
				author: "Good Author",
				copies: 10,
				price:  3000,
			},
			wantBook: &Book{
				Title:  "Good Title",
				Author: "Good Author",
				Copies: 10,
				Price:  3000,
			},
			wantErr: nil,
		},
		{
			name: "Book without Title field",
			args: args{
				author: "Bad Author",
				copies: 10,
				price:  3000,
			},
			wantBook: nil,
			wantErr:  ErrNewBookValidationFailed,
		},
		{
			name: "Book without Author field",
			args: args{
				title:  "Bad Title",
				copies: 10,
				price:  3000,
			},
			wantBook: nil,
			wantErr:  ErrNewBookValidationFailed,
		},
		{
			name: "Book without Copies field",
			args: args{
				author: "Bad Author",
				title:  "Bad Title",
				price:  3000,
			},
			wantBook: nil,
			wantErr:  ErrNewBookValidationFailed,
		},
		{
			name: "Book without Price field",
			args: args{
				author: "Bad Author",
				title:  "Bad Title",
				copies: 10},
			wantBook: nil,
			wantErr:  ErrNewBookValidationFailed,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBook, err := NewBook(tt.args.title, tt.args.author, tt.args.copies, tt.args.price)
			if err != tt.wantErr {
				t.Errorf("NewBook() err = %v, wantErr: [%v]", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBook, tt.wantBook) {
				t.Errorf("NewBook() = %v, want %v", gotBook, tt.wantBook)
			}
		})
	}
}
