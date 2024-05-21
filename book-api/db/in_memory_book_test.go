package db

import (
	book "github.com/mgll/golab/book-api/model/book"
	"testing"
)

func TestAddElement(t *testing.T) {
	b := book.New("A", "B")

	err := Add(b)
	if err != nil {
		t.Fatalf("Adding an element to in memory database failed")
	}

	TearDownForTest()
}

func TestAddAndRetrieveElement(t *testing.T) {
	b := book.New("A", "B")

	err := Add(b)
	if err != nil {
		t.Fatalf("Adding an element to in memory database failed")
	}

	retrievedBook := Get(1)
	if retrievedBook == nil {
		t.Fatalf("Couldn't get book")
	}

	TearDownForTest()
}

func TestAddAndRetrieveMultipleElements(t *testing.T) {
	b1 := book.New("A", "B")
	err := Add(b1)
	if err != nil {
		t.Fatalf("Adding an element to in memory database failed")
	}

	b2 := book.New("C", "D")
	err = Add(b2)
	if err != nil {
		t.Fatalf("Adding an element to in memory database failed")
	}

	dbSize := Size()
	if dbSize != 2 {
		t.Fatalf("Db size is not 2")
	}

	rb1 := Get(1)
	if rb1 == nil {
		t.Fatalf("Couldn't get book")
	}

	if rb1.Name != "A" || rb1.AuthorName != "B" {
		t.Fatalf("Book name or author nae didn't match")
	}

	rb2 := Get(2)
	if rb2 == nil {
		t.Fatalf("Couldn't get book")
	}

	if rb2.Name != "C" || rb2.AuthorName != "D" {
		t.Fatalf("Book name or author nae didn't match")
	}

	TearDownForTest()
}
