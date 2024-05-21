package dto

import "testing"

func TestNewCreateDto(t *testing.T) {
	name := "First book"
	authorName := "John Doe"
	createDto := NewCreateDto(name, authorName)

	if createDto == nil {
		t.Fatalf("createDto is nil")
	}

	if createDto.name != name {
		t.Fatalf("Name of DTO doesn't match %s", name)
	}

	if createDto.authorName != authorName {
		t.Fatalf("Author name of DTO doesn't match %s", authorName)
	}
}
