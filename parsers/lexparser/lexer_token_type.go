package lexparser

// LexTokenType is an implementation of types.TokenType
// Is an extension of type int
type LexTokenType int

const (
	rangePrecedence      int = 1 << iota
	groupingPrecedence   int = 1 << iota
	groupbracePrecedence int = 1 << iota
	concatPrecedence     int = 1 << iota
	orPrecedence         int = 1 << iota
	prefixPrecedence     int = 1 << iota
	postfixPrecedence    int = 1 << iota
)

// Leftparen
// Rightparen
// Leftbrace
// Rightbrace
// Pipe
// Middledot
// Minus
// Plus
// Asterisk
// Question
// Carret
// Char
// EOF
const (
	Lefparen   LexTokenType = 1 << iota
	Rightparen LexTokenType = 1 << iota
	Leftbrace  LexTokenType = 1 << iota
	Rightbrace LexTokenType = 1 << iota
	Pipe       LexTokenType = 1 << iota
	Middledot  LexTokenType = 1 << iota
	Minus      LexTokenType = 1 << iota
	Plus       LexTokenType = 1 << iota
	Asterisk   LexTokenType = 1 << iota
	Question   LexTokenType = 1 << iota
	Carret     LexTokenType = 1 << iota
	Char       LexTokenType = 1 << iota
	EOF        LexTokenType = 1 << iota
)

// GetPrecedence returns the precedence of the LexTokenType l
func (l LexTokenType) GetPrecedence() int {
	switch l {
	case Carret:
		return prefixPrecedence
	case Plus:
		fallthrough
	case Question:
		fallthrough
	case Asterisk:
		return postfixPrecedence
	case Minus:
		return rangePrecedence
	case Pipe:
		return orPrecedence
	case Rightparen:
		fallthrough
	case Lefparen:
		return groupingPrecedence
	case Leftbrace:
		fallthrough
	case Rightbrace:
		return groupbracePrecedence
	case Char:
		return concatPrecedence
	}
	return 0
}

func (l LexTokenType) isPunctuator() bool {
	switch l {
	case Lefparen:
		fallthrough
	case Rightparen:
		fallthrough
	case Leftbrace:
		fallthrough
	case Rightbrace:
		fallthrough
	case Pipe:
		fallthrough
	case Plus:
		fallthrough
	case Minus:
		fallthrough
	case Asterisk:
		fallthrough
	case Question:
		fallthrough
	case Carret:
		return true
	}
	return false
}
