package prattparser

import "github.com/lfmexi/golalr/symbols"

// PrefixParselet is an interface that generates parsers for a prefix operator
type PrefixParselet interface {
	Parse(p *Parser, t symbols.Token) (Expression, error)
}

// InfixParselet is an interface that generates parsers for a prefix operator
type InfixParselet interface {
	Parse(p *Parser, left Expression, t symbols.Token) (Expression, error)
}
