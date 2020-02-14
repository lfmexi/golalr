package grammarparser

import (
	"github.com/lfmexi/golalr/prattparser"
	"github.com/lfmexi/golalr/prattparser/symbols"
)

type simpleParselet struct{}

func (s simpleParselet) Parse(p *prattparser.Parser, t symbols.Token) (prattparser.Expression, error) {
	return newSimpleExpression(t), nil
}
