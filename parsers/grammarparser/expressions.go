package grammarparser

import "github.com/lfmexi/golalr/prattparser"
import "github.com/lfmexi/golalr/symbols"
import "github.com/lfmexi/golalr/expressions"

type ProductionExpression struct {
	next  prattparser.Expression
	token symbols.Token
}

func (p *ProductionExpression) SetToken(token symbols.Token) {
	p.token = token
}

func (p ProductionExpression) Token() symbols.Token {
	return p.token
}

func (p ProductionExpression) Next() expressions.SimpleExpression {
	next := p.next
	if next == nil {
		return nil
	}
	return p.next.(*ProductionExpression)
}

func (e *ProductionExpression) setNext(n expressions.SimpleExpression) {
	e.next = n
}

func newSimpleExpression(t symbols.Token) prattparser.Expression {
	return &ProductionExpression{nil, t}
}
