package grammarparser

import (
	"github.com/lfmexi/golalr/expressions"
	"github.com/lfmexi/golalr/prattparser"
	"github.com/lfmexi/golalr/symbols"
)

// ProductionParser is a parser for left recursive grammars
type ProductionParser struct {
	innerParser *prattparser.Parser
}

// NewProductionParser creates a new Production parser
func NewProductionParser(scanner symbols.TokenIterator) ProductionParser {
	parser := prattparser.NewParser(scanner)
	parser.RegisterPrefixParselet(SymbolID, &simpleParselet{})
	parser.RegisterInfixParselet(SymbolID, &idInfixParselet{})
	return ProductionParser{&parser}
}

// Parse initiates the parsing process
func (p *ProductionParser) Parse() (expressions.SimpleExpression, error) {
	expression, err := p.innerParser.Parse(0)

	if err != nil {
		return nil, err
	}

	productionExpression := expression.(*ProductionExpression)
	return productionExpression, nil
}
