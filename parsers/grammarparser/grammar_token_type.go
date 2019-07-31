package grammarparser

// ProductionTokenType is the string alias that represents the type of the grammar productions
type ProductionTokenType string

const (
	idPrecedence int = 1 << iota
)

const (
	// SymbolID is the reserved ProductionTokenType for ids
	SymbolID ProductionTokenType = "id"
	// EOF is the reserved ProductionTokenType for end of file
	EOF ProductionTokenType = "EOF"
)

// GetPrecedence returns the Precedence of the given tokenType
func (p ProductionTokenType) GetPrecedence() int {
	switch p {
	case SymbolID:
		return idPrecedence
	}
	return 0
}
