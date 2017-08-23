package lexer

import (
	"fmt"

	"github.com/lfmexi/golalr/expressions"
	"github.com/lfmexi/golalr/parsers/lexparser"
)

func expandExpression(exp *lexparser.LexerExpression) (*lexparser.LexerExpression, error) {
	newExp, err := expandRangeNode(exp)
	if err != nil {
		return nil, err
	}
	newLexerExpression := newExp.(lexparser.LexerExpression)
	return &newLexerExpression, nil
}

func expandRangeNode(n *lexparser.LexerExpression) (expressions.BinaryExpression, error) {
	if n.IsLeaf() {
		return *n, nil
	}

	if n.Token().TokenType() != lexparser.Minus && n.Token().TokenType() != lexparser.Minus {
		if n.Left() != nil {
			var err error
			leftExpression := n.Left().(lexparser.LexerExpression)
			newLeft, err := expandRangeNode(&leftExpression)
			if err != nil {
				return nil, err
			}
			n.SetLeft(newLeft)
		}

		if n.Right() != nil {
			var err error
			rightExpression := n.Right().(lexparser.LexerExpression)
			newRight, err := expandRangeNode(&rightExpression)
			if err != nil {
				return nil, err
			}
			n.SetRight(newRight)
		}
		return *n, nil
	}

	newNode, err := createOrAST(n)

	if err != nil {
		return nil, err
	}

	return newNode, nil
}

func createOrAST(n *lexparser.LexerExpression) (expressions.BinaryExpression, error) {
	leftToken := n.Left().(lexparser.LexerExpression).Token()
	rightToken := n.Right().(lexparser.LexerExpression).Token()

	leftBytes := []byte(leftToken.Text())
	leftRune := rune(leftBytes[0])

	rightBytes := []byte(rightToken.Text())
	rightRune := rune(rightBytes[0])

	if leftRune > rightRune {
		return nil, fmt.Errorf("Cannot do a range from %s to %s", string(leftRune), string(rightRune))
	}

	return getOrExpression(leftRune, rightRune)
}

func getOrExpression(from rune, to rune) (expressions.BinaryExpression, error) {
	accum := string(from)
	for index := from + 1; index <= to; index++ {
		accum = fmt.Sprintf("%s|%s", accum, string(index))
	}
	scanner := lexparser.NewLexScanner(accum)
	parser := lexparser.NewLexerParser(scanner)
	expression, _ := parser.GetInnerParser().Parse(0)

	return expression.(lexparser.LexerExpression), nil
}
