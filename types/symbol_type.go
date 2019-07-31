package types

// SymbolType is the string alias for Symbols
type SymbolType string

// GetPrecedence returns the precedence of the SymbolType
func (SymbolType) GetPrecedence() int {
	return 0
}

const (
	// EOF is the end of file Symbol
	EOF SymbolType = "EOF"
	// UnknownSymbol represents any unknown symbol for the grammar
	UnknownSymbol SymbolType = "UNKNOWN_SYMBOL"
)
