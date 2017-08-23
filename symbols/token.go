package symbols

import "github.com/lfmexi/golalr/types"

// Token is the interface that will contain the TokenType of an accepted character of
// a previously defined language
// TokenType is the TokenType
// Text string is the recognized symbol
type Token interface {
	TokenType() types.TokenType
	Text() string
}
