package lexparser

import (
	"reflect"
	"testing"

	"github.com/lfmexi/golalr/prattparser"
)

func TestLexScanner_Next(t *testing.T) {
	lexScanner := NewLexScanner("125ewd*(a|nb)")
	tests := []struct {
		name string
		l    prattparser.TokenIterator
		want string
	}{
		{"1", lexScanner, "1"},
		{"2", lexScanner, "2"},
		{"5", lexScanner, "5"},
		{"e", lexScanner, "e"},
		{"w", lexScanner, "w"},
		{"d", lexScanner, "d"},
		{"*", lexScanner, "*"},
		{"(", lexScanner, "("},
		{"a", lexScanner, "a"},
		{"|", lexScanner, "|"},
		{"n", lexScanner, "n"},
		{"b", lexScanner, "b"},
		{")", lexScanner, ")"},
		{"EOF", lexScanner, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Next().Text(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LexScanner.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
