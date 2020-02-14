package golalr

import (
	"github.com/lfmexi/golalr/lexer"
	"github.com/lfmexi/golalr/prattparser/types"
)

type Symbol interface {
}

type lalrParser struct{}

func newLalr() *lalrParser {
	return &lalrParser{}
}

type GolalrConfig struct {
	lexer.Builder
	parser *lalrParser
}

func NewGolalrConfig() *GolalrConfig {
	return &GolalrConfig{
		lexer.NewLexerBuilder(),
		nil,
	}
}

func (g *GolalrConfig) AddTerminalDefinition(id string, regExpression string, action lexer.Action) {
	// TODO - Add the terminal definition to parser builder
	g.Builder.AddTerminalDefinition(types.SymbolType(id), regExpression, action)
}

func (*GolalrConfig) AddNonTerminal(id string) {

}

func (*GolalrConfig) AddProduction(dest string, prod string, action func(map[string]string)) {

}

func (*GolalrConfig) CreateParser() (*Parser, error) {
	parser := newParser()
	return parser, nil
}
