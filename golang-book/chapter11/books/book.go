package books

import (
	"encoding/json"
	"strings"
)

type Book struct {
	Title   string
	Author	string
	Pages	int
}

func (b *Book) CategoryByLength() string {
	if b.Pages > 300 {
		return "NOVEL"
	}

	return "SHORT STORY"
}
func (b *Book) AuthorLastName() string {
	author := strings.SplitN(b.Author, " ", 10)
	return author[len(author) - 1]
}

func NewBookFromJSON(jsn string) (Book, error) {
	var book Book
	err := json.Unmarshal([]byte(jsn), &book)
	return book, err
}