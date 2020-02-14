package lexparser

import (
	"github.com/lfmexi/golalr/prattparser"
	"github.com/lfmexi/golalr/prattparser/symbols"
)

func newSimpleExpression(t symbols.Token) prattparser.Expression {
	return LexerExpression{nil, nil, t}
}

func newPrefixExpression(t symbols.Token, l prattparser.Expression) prattparser.Expression {
	return LexerExpression{
		nil,
		l,
		t,
	}
}

func newOperatorExpression(left prattparser.Expression, right prattparser.Expression, t symbols.Token) prattparser.Expression {
	return LexerExpression{
		left,
		right,
		t,
	}
}

func newPostfixExpression(left prattparser.Expression, t symbols.Token) prattparser.Expression {
	return LexerExpression{
		left,
		nil,
		t,
	}
}
