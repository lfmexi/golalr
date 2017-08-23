package lexparser

import (
	"fmt"

	"github.com/lfmexi/golalr/expressions"
	"github.com/lfmexi/golalr/prattparser"
	"github.com/lfmexi/golalr/symbols"
)

// LexerExpression implements the prattparser.Expression interface
// Its structure builds an AST Node
type LexerExpression struct {
	left  prattparser.Expression
	right prattparser.Expression
	token symbols.Token
}

// SetToken sets the token that will represent this expression
func (l *LexerExpression) SetToken(token symbols.Token) {
	l.token = token
}

// Token gets the symbols.Token of the expression
func (l LexerExpression) Token() symbols.Token {
	return l.token
}

// Left gets the left expression
func (l LexerExpression) Left() expressions.BinaryExpression {
	left := l.left
	if left == nil {
		return nil
	}
	return l.left.(LexerExpression)
}

// SetLeft is a setter for the left node
func (l *LexerExpression) SetLeft(e expressions.BinaryExpression) {
	l.left = e
}

// Right gets the right expression
func (l LexerExpression) Right() expressions.BinaryExpression {
	right := l.right
	if right == nil {
		return nil
	}
	return l.right.(LexerExpression)
}

// SetRight is a setter for the right node
func (l *LexerExpression) SetRight(e expressions.BinaryExpression) {
	l.right = e
}

// IsLeaf indicates if the current LexerExpression is a leaf node
func (l LexerExpression) IsLeaf() bool {
	return l.Left() == nil && l.Right() == nil
}

// ToString returns a JSON-like structure of the LexerExpression
func (l LexerExpression) ToString() string {
	text := l.Token().Text()

	if l.IsLeaf() {
		return fmt.Sprintf("\"%s\"", text)
	}

	var left, right string

	if l.Left() != nil {
		left = "\"left\":" + l.Left().(LexerExpression).ToString()
	}

	if l.Right() != nil {
		if l.Left() != nil {
			right = ","
		}
		right = right + "\"right\":" + l.Right().(LexerExpression).ToString()
	}
	return fmt.Sprintf("{\"%s\": {%s%s}}", text, left, right)
}
