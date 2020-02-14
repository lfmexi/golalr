package grammarparser

import "github.com/lfmexi/golalr/prattparser"
import "github.com/lfmexi/golalr/prattparser/symbols"

type idInfixParselet struct{}

func (idInfixParselet) Parse(p *prattparser.Parser, previous prattparser.Expression, t symbols.Token) (prattparser.Expression, error) {
	nextExpression := newSimpleExpression(t).(*ProductionExpression)
	token := p.LookNext()

	previousExp := previous.(*ProductionExpression)
	previousExp.setNext(nextExpression)

	if token.TokenType().GetPrecedence() < t.TokenType().GetPrecedence() {
		return previous, nil
	}

	infix := p.GetInfixParselet(token.TokenType())

	token, _ = p.Consume(nil)

	_, err := infix.Parse(p, nextExpression, token)

	if err != nil {
		return nil, err
	}

	return previous, nil
}
