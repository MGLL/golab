package model

type Book struct {
	id         int
	Name       string
	AuthorName string
}

func New(name, authorName string) *Book {
	return &Book{
		Name: name, AuthorName: authorName,
	}
}

func (b *Book) SetId(id int) {
	b.id = id
}

func (b *Book) GetId() int {
	return b.id
}
