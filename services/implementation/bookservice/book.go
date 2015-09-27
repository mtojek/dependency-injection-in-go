package bookservice

import "github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"

type book struct {
	title string
}

func newBook(title string) shared.Book {
	return &book{title: title}
}

// Name method returns book title.
func (b *book) Title() string {
	return b.title
}
