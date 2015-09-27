# Dependency Injection in Go

This project shows how to solve (DIY) a dependency injection problem in Go.

Keywords: dependency injection, inject, injector, DIY

## Introduction

Dependency injection seems to be a key factor to write a DRI code, especially when a developer does not want to care about passing a service through few layers, always accompanied by a long list of constructor arguments. This boilerplate usually makes the code unreadable and unnecessarily longer than the crucial business logic.

The way I figured this out, is presented in a "proof of concept" simple project. I tried to eliminate as much of not related to DI source code, thus sometimes you may wonder if this simplicity really needs a DI pattern.

## User Stories

* There is a library in which a user can borrow a book.
* The user has name
* The book has a title
* The user can borrow a book.

## Design of the Solution

The application has a separated "services" package, which contains two another packages - "implementation" and "interfaces". The first one contains several application services (also organized in packages) like a `BookService`, `BorrowService`, `UserService`, `LoggerService` and `BorrowingFormatter`.

I assume that all mentioned services should be stateless, thus they can be represented by only one instance per service. The separation of interfaces and implementation prevents from problematic dependency cycles.

### BookService - Sample Service

Package: `bookservice`

The initialization part has been separated because of SRP (Single Responsibility Pattern). All initialization operations have been moved to a file called `book_service_init.go`. When a client service requests an instance of `BookService` it should call the `bookservice.Instance()` method. The instance will lazily created if necessary:

~~~ go
var (
	once     sync.Once
	instance interfaces.BookService
)

func newBookService() interfaces.BookService {
	return &bookService{
		loggerService: loggerservice.Instance(),
	}
}

// Instance method returns a singleton instance of BookService.
func Instance() interfaces.BookService {
	once.Do(func() {
		instance = newBookService()
	})
	return instance
}
~~~

The implementation of `BookService` has been placed in `book_service.go` and contains only business logic and structure representation:

~~~ go
type bookService struct {
	loggerService interfaces.LoggerService
}

func (b *bookService) CreateBook(title string) shared.Book {
	b.loggerService.Info("New book created: %v", title)
	return newBook(title)
}
~~~

One may say that having interfaces causes writing more code. Actually it supports testing, because it allows mocking services:

~~~ go
func TestCreateBook(t *testing.T) {
	assert := assert.New(t)

	// given
	logger := new(mockedLogger)
	sut := &bookService{loggerService: logger}

	// when
	sut.CreateBook("book")

	// then
	assert.True(logger.invoked, "LoggerService's Info method should be invoked")
}
~~~


