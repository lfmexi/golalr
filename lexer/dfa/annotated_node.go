package dfa

import (
	"github.com/lfmexi/golalr/parsers/lexparser"
	"github.com/lfmexi/golalr/symbols"
)

type annotatedASTNode struct {
	token          symbols.Token
	firstPositions []int
	lastPositions  []int
	nextPositions  []int
	left           *annotatedASTNode
	right          *annotatedASTNode
}

func (n *annotatedASTNode) isNullable() bool {
	if n.token.TokenType() == lexparser.Pipe {
		return n.left.isNullable() || n.right.isNullable()
	}

	if n.token.TokenType() == lexparser.Middledot {
		return n.left.isNullable() && n.right.isNullable()
	}

	return n.token.TokenType() == lexparser.Asterisk || n.token.TokenType() == lexparser.Question || n.token.TokenType() == lexparser.EOF
}

func (n *annotatedASTNode) setFirstPositions() {
	if n.left == nil && n.right == nil {
		return
	}

	leftFirsts := make([]int, 0)
	rightFirsts := make([]int, 0)

	if n.left != nil {
		n.left.setFirstPositions()
		leftFirsts = n.left.firstPositions
	}

	if n.right != nil {
		n.right.setFirstPositions()
		rightFirsts = n.right.firstPositions
	}

	switch n.token.TokenType() {
	case lexparser.Pipe:
		n.firstPositions = append(leftFirsts, rightFirsts...)
	case lexparser.Middledot:
		if n.left.isNullable() {
			n.firstPositions = append(leftFirsts, rightFirsts...)
		} else {
			n.firstPositions = append(make([]int, 0), leftFirsts...)
		}
	case lexparser.Question:
		fallthrough
	case lexparser.Plus:
		fallthrough
	case lexparser.Asterisk:
		n.firstPositions = append(make([]int, 0), leftFirsts...)
	}
}

func (n *annotatedASTNode) setLastPositions() {
	if n.left == nil && n.right == nil {
		return
	}

	leftLasts := make([]int, 0)
	rightLasts := make([]int, 0)

	if n.left != nil {
		n.left.setLastPositions()
		leftLasts = n.left.lastPositions
	}

	if n.right != nil {
		n.right.setLastPositions()
		rightLasts = n.right.lastPositions
	}

	switch n.token.TokenType() {
	case lexparser.Pipe:
		n.lastPositions = append(leftLasts, rightLasts...)
	case lexparser.Middledot:
		if n.right.isNullable() {
			n.lastPositions = append(leftLasts, rightLasts...)
		} else {
			n.lastPositions = append(make([]int, 0), rightLasts...)
		}
	case lexparser.Question:
		fallthrough
	case lexparser.Plus:
		fallthrough
	case lexparser.Asterisk:
		n.lastPositions = append(make([]int, 0), leftLasts...)
	}
}

func contains(slice []int, value int) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}
	return false
}

func (n *annotatedASTNode) setNextPositions(where []int, set []int) {
	if n.left == nil && n.right == nil {
		if contains(where, n.firstPositions[0]) {
			n.nextPositions = append(n.nextPositions, set...)
		}
		return
	}

	switch n.token.TokenType() {
	case lexparser.Middledot:
		left := n.left
		right := n.right
		left.setNextPositions(left.lastPositions, right.firstPositions)
		if right.isNullable() && right.token.TokenType() != lexparser.EOF {
			left.setNextPositions(left.lastPositions, set)
		}
		right.setNextPositions(where, set)
	case lexparser.Plus:
		fallthrough
	case lexparser.Asterisk:
		left := n.left
		newSet := append(n.firstPositions, set...)
		left.setNextPositions(n.lastPositions, newSet)
	case lexparser.Question:
		left := n.left
		left.setNextPositions(n.lastPositions, set)
	case lexparser.Pipe:
		left := n.left
		right := n.right
		left.setNextPositions(where, set)
		right.setNextPositions(where, set)
	}
}
