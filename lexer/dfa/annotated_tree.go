package dfa

import (
	"fmt"

	"github.com/lfmexi/golalr/expressions"
	"github.com/lfmexi/golalr/parsers/lexparser"
	"github.com/lfmexi/golalr/types"
)

type innerToken struct {
	text  string
	ttype types.TokenType
}

func (i innerToken) Text() string {
	return i.text
}

func (i innerToken) TokenType() types.TokenType {
	return i.ttype
}

// AnnotatedAST represents a LexerExpression in a tree
// with annotated nodes. It contains metadata that will help
// in the creation of an DFA
type AnnotatedAST struct {
	currentPosition int
	root            *annotatedASTNode
	leafs           map[int]*annotatedASTNode
	symbols         map[string]map[int]int
}

// NewAnnotatedAST creates a new instance of AnnotatedAST using a
// lexparser.LexerExpression for creating annotated nodes
func NewAnnotatedAST(exp expressions.BinaryExpression) *AnnotatedAST {
	a := &AnnotatedAST{
		0,
		nil,
		make(map[int]*annotatedASTNode),
		make(map[string]map[int]int),
	}
	node := a.annotateExpression(exp)
	concatNode := &annotatedASTNode{
		&innerToken{"${concat}", lexparser.Middledot},
		make([]int, 0),
		make([]int, 0),
		make([]int, 0),
		node,
		a.getEOFNode(),
	}
	if a.root == nil {
		a.root = concatNode
	}
	return a
}

func (a *AnnotatedAST) getEOFNode() *annotatedASTNode {
	a.currentPosition++
	eofNode := &annotatedASTNode{
		&innerToken{"${EOF}", lexparser.EOF},
		[]int{a.currentPosition},
		[]int{a.currentPosition},
		make([]int, 0),
		nil,
		nil,
	}
	a.leafs[a.currentPosition] = eofNode
	a.symbols["EOF"] = make(map[int]int)
	a.symbols["EOF"][a.currentPosition] = a.currentPosition
	return eofNode
}

func (a *AnnotatedAST) annotateExpression(exp expressions.BinaryExpression) *annotatedASTNode {
	if exp.IsLeaf() {
		a.currentPosition++
		node := &annotatedASTNode{
			exp.Token(),
			[]int{a.currentPosition},
			[]int{a.currentPosition},
			make([]int, 0),
			nil,
			nil,
		}
		if a.symbols[exp.Token().Text()] == nil {
			a.symbols[exp.Token().Text()] = make(map[int]int)
		}
		a.symbols[exp.Token().Text()][a.currentPosition] = a.currentPosition
		a.leafs[a.currentPosition] = node
		return node
	}

	node := &annotatedASTNode{
		exp.Token(),
		make([]int, 0),
		make([]int, 0),
		make([]int, 0),
		nil,
		nil,
	}

	if exp.Left() != nil {
		leftExpression := exp.Left()
		node.left = a.annotateExpression(leftExpression)
	}

	if exp.Right() != nil && exp.Token().TokenType() != lexparser.Carret {
		rightExpression := exp.Right()
		node.right = a.annotateExpression(rightExpression)
	}

	if exp.Token().TokenType() == lexparser.Carret {
		a.currentPosition++
		rightText := exp.Right().Token().Text()
		newText := fmt.Sprintf("^%v", rightText)
		node.token = &innerToken{newText, lexparser.Char}
		node.firstPositions = []int{a.currentPosition}
		node.lastPositions = []int{a.currentPosition}
		if a.symbols[newText] == nil {
			a.symbols[newText] = make(map[int]int)
		}
		a.symbols[newText][a.currentPosition] = a.currentPosition
		a.leafs[a.currentPosition] = node
	}

	return node
}
