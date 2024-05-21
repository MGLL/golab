package model

import "testing"

func TestCreateBook(t *testing.T) {
	name := "First book"
	authorName := "John Doe"
	book := New(name, authorName)

	if book == nil {
		t.Fatalf("book is nil")
	}

	if book.Name != name {
		t.Fatalf("Name of book doesn't match %s", name)
	}

	if book.AuthorName != authorName {
		t.Fatalf("Author name of book doesn't match %s", authorName)
	}
}
