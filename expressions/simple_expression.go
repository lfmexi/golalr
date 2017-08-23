package expressions

import (
	"github.com/lfmexi/golalr/symbols"
)

type SimpleExpression interface {
	Token() symbols.Token
	Next() SimpleExpression
}
