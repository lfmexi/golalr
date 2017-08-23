package lexparser

import "github.com/lfmexi/golalr/types"

type LexToken struct {
	text  string
	tType types.TokenType
}

func (l LexToken) Text() string {
	return l.text
}

func (l LexToken) TokenType() types.TokenType {
	return l.tType
}
