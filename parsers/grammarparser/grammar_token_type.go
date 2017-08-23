package grammarparser

type ProductionTokenType string

const (
	idPrecedence int = 1 << iota
)

const (
	SymbolID ProductionTokenType = "id"
	EOF      ProductionTokenType = "EOF"
)

func (p ProductionTokenType) GetPrecedence() int {
	switch p {
	case SymbolID:
		return idPrecedence
	}
	return 0
}
