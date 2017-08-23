package lexparser

import (
	"github.com/lfmexi/golalr/prattparser"
	"github.com/lfmexi/golalr/symbols"
)

type simpleParselet struct{}

func (s simpleParselet) Parse(p *prattparser.Parser, t symbols.Token) (prattparser.Expression, error) {
	expression := newSimpleExpression(t)
	return expression, nil
}

type rangeParselet struct{}

func (r rangeParselet) Parse(p *prattparser.Parser, t symbols.Token) (prattparser.Expression, error) {

	leftToken, err := p.Consume(Char)
	if err != nil {
		return nil, err
	}

	minustoken, err := p.Consume(Minus)
	if err != nil {
		return nil, err
	}

	rightToken, err := p.Consume(Char)

	if err != nil {
		return nil, err
	}

	if _, err := p.Consume(Rightbrace); err != nil {
		return nil, err
	}

	expression := newOperatorExpression(newSimpleExpression(leftToken), newSimpleExpression(rightToken), minustoken)

	return expression, nil
}

type prefixOperatorParselet struct {
}

func (o prefixOperatorParselet) Parse(p *prattparser.Parser, t symbols.Token) (prattparser.Expression, error) {
	token, err := p.Consume(Char)

	if err != nil {
		return nil, err
	}

	right := newSimpleExpression(token)
	expression := newPrefixExpression(t, right)
	return expression, nil
}

type groupParselet struct{}

func (r groupParselet) Parse(p *prattparser.Parser, t symbols.Token) (prattparser.Expression, error) {
	expression, err := p.Parse(t.TokenType().GetPrecedence())

	if err != nil {
		return nil, err
	}

	if _, err := p.Consume(Rightparen); err != nil {
		return nil, err
	}

	return expression, nil
}
