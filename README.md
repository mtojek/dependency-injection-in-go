# Dependency Injection in Go

[![Build Status](https://travis-ci.org/mtojek/dependency-injection-in-go.svg?branch=master)](https://travis-ci.org/mtojek/dependency-injection-in-go)

This project shows how to solve (DIY) a dependency injection problem in Go.

Status: **Done**

Keywords: dependency injection, inject, injector, DIY

## Introduction

Dependency injection seems to be a key factor while writing a DRY code, especially when a developer does not want to care about passing a service through few layers, always accompanied by a long list of constructor arguments. This boilerplate usually makes the code unreadable and unnecessarily longer than the crucial business logic.

The way I figured this out, is presented in a "proof of concept" simple project. The application bases on a graph provided by [facebookgo/inject](https://github.com/facebookgo/inject), which is responsible for resolving dependencies between services.

## User stories

* There is a library in which a user can borrow a book.
* The user has name.
* The book has a title.
* The user can borrow a book.

## Design of the solution

The application has a separated "services" package, which contains only two packages - "implementation" and "interfaces". The first one contains several application services (also organized in packages) like a `BookService`, `BorrowService`, `UserService`, `LoggerService` and `BorrowingFormatter`.

I assume that all mentioned services should be stateless, thus they can be represented by only one instance per service. The separation of interfaces and implementation prevents from problematic dependency cycles and makes easier services mocking.

### BookService - Sample Service

As there can be noticed, a logger service will be injected by the graph population (description below). No constructors, special initialization procedures have been used.

~~~ go
package bookservice

import ...

// BookService allows to create books.
type BookService struct {
	LoggerService interfaces.LoggerService `inject:""`
}

// CreateBook method is responsible for creating new book.
func (b *BookService) CreateBook(title string) shared.Book {
	b.LoggerService.Info("New book created: %v", title)
	return newBook(title)
}

~~~

If we use interfaces for injections, there will be a possibility to write some tests, including mocking them externally (exported fields).

~~~ go
func TestCreateBook(t *testing.T) {
	assert := assert.New(t)

	// given
	logger := new(mockedLogger)
	sut := &BookService{logger}

	// when
	sut.CreateBook("book")

	// then
	assert.True(logger.invoked, "LoggerService's Info method should be invoked")
}
~~~

### BorrowingFormatter - new instance

In some situations it is useful to always inject a new instance of service. If a service wants to inject a new instance of `BorrowingFormatter` it should call `borrowingformatter.New()` method:

~~~ go
// Borrow method is responsible for borrowing book by user.
func (b *borrowService) Borrow(user shared.User, book shared.Book) {
	formatter := borrowingformatter.New()
	formatted := formatter.Format(user, book)

	b.loggerService.Info(formatted)
}
~~~

Hint: for testing purposes I recommend to prepare a "new instance provider", which can produce any mocked instance we want. The provider pattern is not applied to this project.

