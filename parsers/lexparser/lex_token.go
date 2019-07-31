package lexparser

import "github.com/lfmexi/golalr/types"

// LexToken is a Token implementation for the lexparser
type LexToken struct {
	text  string
	tType types.TokenType
}

// Text retrieves the text contained in the Token
func (l LexToken) Text() string {
	return l.text
}

// TokenType gives the type of the Token
func (l LexToken) TokenType() types.TokenType {
	return l.tType
}
