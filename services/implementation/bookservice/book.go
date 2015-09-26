package bookservice

import "github.com/mtojek/dependency-injection-in-go/services/interfaces/shared"

type book struct {
	name string
}

func newBook(name string) shared.Book {
	return &book{name: name}
}

// Name method returns book name.
func (b *book) Name() string {
	return b.name
}
