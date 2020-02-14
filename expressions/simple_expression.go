package expressions

import (
	"github.com/lfmexi/golalr/prattparser/symbols"
)

type SimpleExpression interface {
	Token() symbols.Token
	Next() SimpleExpression
}
