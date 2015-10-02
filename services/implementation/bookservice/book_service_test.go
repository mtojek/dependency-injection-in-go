package bookservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
