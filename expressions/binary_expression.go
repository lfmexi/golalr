package expressions

import "github.com/lfmexi/golalr/symbols"

// BinaryExpression is an interface that tries to represent all the
// expresions with two childs
type BinaryExpression interface {
	IsLeaf() bool
	Token() symbols.Token
	Left() BinaryExpression
	Right() BinaryExpression
}
