package dto

type BookCreate struct {
	name       string
	authorName string
}

func NewCreateDto(name, authorName string) *BookCreate {
	return &BookCreate{name, authorName}
}
