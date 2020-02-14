package grammarparser

import "github.com/lfmexi/golalr/prattparser"
import "github.com/lfmexi/golalr/prattparser/symbols"
import "github.com/lfmexi/golalr/expressions"

// ProductionExpression is the defined expression for grammar productions
type ProductionExpression struct {
	next  prattparser.Expression
	token symbols.Token
}

// SetToken sets the token to the expression
func (p *ProductionExpression) SetToken(token symbols.Token) {
	p.token = token
}

// Token retrieves the token from the expression
func (p ProductionExpression) Token() symbols.Token {
	return p.token
}

// Next gets the next expression
func (p ProductionExpression) Next() expressions.SimpleExpression {
	next := p.next
	if next == nil {
		return nil
	}
	return p.next.(*ProductionExpression)
}

func (p *ProductionExpression) setNext(n expressions.SimpleExpression) {
	p.next = n
}

func newSimpleExpression(t symbols.Token) prattparser.Expression {
	return &ProductionExpression{nil, t}
}
