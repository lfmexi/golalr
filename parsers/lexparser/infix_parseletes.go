package lexparser

import (
	"github.com/lfmexi/golalr/prattparser"
	"github.com/lfmexi/golalr/prattparser/symbols"
)

type charInfixParselet struct{}

func (charInfixParselet) Parse(p *prattparser.Parser, left prattparser.Expression, t symbols.Token) (prattparser.Expression, error) {
	charExpression := newSimpleExpression(t)
	token := p.LookNext()
	concatToken := LexToken{"${concat}", Middledot}

	if token.TokenType().GetPrecedence() <= t.TokenType().GetPrecedence() {
		concatExpression := newOperatorExpression(left, charExpression, concatToken)
		return concatExpression, nil
	}

	infix := p.GetInfixParselet(token.TokenType())

	token, _ = p.Consume(nil)

	right, err := infix.Parse(p, charExpression, token)

	if err != nil {
		return nil, err
	}
	return newOperatorExpression(left, right, concatToken), nil
}

type binaryOperatorParselet struct{}

func (binaryOperatorParselet) Parse(p *prattparser.Parser, left prattparser.Expression, t symbols.Token) (prattparser.Expression, error) {
	right, err := p.Parse(t.TokenType().GetPrecedence())

	if err != nil {
		return nil, err
	}

	return newOperatorExpression(left, right, t), nil
}

type rightComplexInfixParselet struct{}

func getComplexRightExpression(p *prattparser.Parser, t symbols.Token) (prattparser.Expression, error) {
	var right prattparser.Expression
	var err error
	if t.TokenType() != Lefparen {
		parselet := p.GetPrefixParselet(t.TokenType())
		right, err = parselet.Parse(p, t)
		if err != nil {
			return nil, err
		}
		return right, nil
	}

	right, err = p.Parse(t.TokenType().GetPrecedence())

	if err != nil {
		return nil, err
	}

	if _, err := p.Consume(Rightparen); err != nil {
		return nil, err
	}
	return right, nil
}

func (rightComplexInfixParselet) Parse(p *prattparser.Parser, left prattparser.Expression, t symbols.Token) (prattparser.Expression, error) {
	right, err := getComplexRightExpression(p, t)

	if err != nil {
		return nil, err
	}

	tokenConcat := LexToken{"${concat}", Middledot}
	token := p.LookNext()

	if token.TokenType().GetPrecedence() <= t.TokenType().GetPrecedence() {
		return newOperatorExpression(left, right, tokenConcat), nil
	}

	token, _ = p.Consume(nil)

	infix := p.GetInfixParselet(token.TokenType())

	newRight, err := infix.Parse(p, right, token)

	if err != nil {
		return nil, err
	}

	return newOperatorExpression(left, newRight, tokenConcat), nil
}

type postfixOperatorParselet struct {
}

func (p postfixOperatorParselet) Parse(parser *prattparser.Parser, left prattparser.Expression, t symbols.Token) (prattparser.Expression, error) {
	return newPostfixExpression(left, t), nil
}
