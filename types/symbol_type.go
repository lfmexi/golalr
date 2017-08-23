package types

type SymbolType string

func (SymbolType) GetPrecedence() int {
	return 0
}

const (
	EOF           SymbolType = "EOF"
	UnknownSymbol SymbolType = "UNKNOWN_SYMBOL"
)
