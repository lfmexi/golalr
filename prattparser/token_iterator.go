package prattparser

import "github.com/lfmexi/golalr/prattparser/symbols"

// TokenIterator is an Iterator of Tokens
type TokenIterator interface {
	Next() symbols.Token

	AddDelimiter(delimiter string)
}
