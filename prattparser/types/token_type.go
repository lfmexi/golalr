package types

// TokenType represents the types of the tokens defined on a language
type TokenType interface {
	GetPrecedence() int
}
