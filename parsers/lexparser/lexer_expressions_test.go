package lexparser

import (
	"testing"

	"github.com/lfmexi/golalr/prattparser"
)

func TestLexerExpression_IsLeaf(t *testing.T) {
	chartoken := LexToken{"h", Char}
	concattoken := LexToken{"$(concat)", Middledot}
	leftexp := newSimpleExpression(LexToken{"a", Char})
	rightexp := newSimpleExpression(LexToken{"b", Char})
	tests := []struct {
		name string
		l    prattparser.Expression
		want bool
	}{
		{name: "simple", l: newSimpleExpression(chartoken), want: true},
		{name: "composed", l: newOperatorExpression(leftexp, rightexp, concattoken), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.(LexerExpression).IsLeaf(); got != tt.want {
				t.Errorf("LexerExpression.IsLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
