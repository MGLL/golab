package db

import (
	"github.com/mgll/golab/book-api/model/book"
	"slices"
)

// TODO: maybe move it to book directly with book.Persist(), book.Fetch(), book.Count(), book.Delete()

var db = make(map[int]*model.Book)

func Add(b *model.Book) error {
	id := b.GetId()

	if id == 0 {
		id = getMaxId()
	}

	db[id] = b
	return nil
}

func getMaxId() int {
	if len(db) == 0 {
		return 1
	}

	var keys []int
	for k := range db {
		keys = append(keys, k)
	}

	return slices.Max(keys) + 1
}

func Get(id int) *model.Book {
	return db[id]
}

func Size() int {
	return len(db)
}

func TearDownForTest() {
	clear(db)
}
