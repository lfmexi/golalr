package symbols

// TokenIterator is an Iterator of Tokens
type TokenIterator interface {
	Next() Token
}
