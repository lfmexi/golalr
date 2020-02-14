package dfa

import (
	"testing"

	"github.com/lfmexi/golalr/parsers/lexparser"
	"github.com/lfmexi/golalr/prattparser/types"
)

func TestDFA_GetTerminal(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		dfa  *DFA
		args args
		want types.TokenType
	}{
		{"(a|b)*abb", BuildDFA(types.SymbolType("sample"), createAnnotatedTree("(a|b)*abb")), args{"abbaabb"}, types.SymbolType("sample")},
		{"id", BuildDFA(types.SymbolType("id"), createAnnotatedTree("((a|b|c|d|e|f)+|_+)((a|b|c|d|e|f)|(0|1|2|3|4))*")), args{"g"}, types.SymbolType("id")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.dfa
			if got := d.GetTerminal(); got != tt.want {
				t.Errorf("DFA.GetTerminal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createAnnotatedTree(regexp string) *AnnotatedAST {
	scanner := lexparser.NewLexScanner(regexp)
	parser := lexparser.NewLexerParser(scanner)
	expression, _ := parser.Parse()
	return NewAnnotatedAST(expression)
}

func TestDFA_Evaluate(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		dfa  *DFA
		args args
		want bool
	}{
		{"(a|b)*abb", BuildDFA(types.SymbolType("sample"), createAnnotatedTree("(a|b)*abb")), args{"abbaabb"}, true},
		{"(a|b)*abb", BuildDFA(types.SymbolType("sample"), createAnnotatedTree("(a|b)*abb")), args{"abbaa"}, false},
		{"(0|1|2|3|4|5|6|7|8|9)+", BuildDFA(types.SymbolType("integer"), createAnnotatedTree("(0|1|2|3|4|5|6|7|8|9)+")), args{"1234"}, true},
		{"(0|1|2|3|4|5|6|7|8|9)+", BuildDFA(types.SymbolType("integer"), createAnnotatedTree("(0|1|2|3|4|5|6|7|8|9)+")), args{"a12"}, false},
		{"ab*c?", BuildDFA(types.SymbolType("rare"), createAnnotatedTree("ab*c?")), args{"abc"}, true},
		{"ab*c?", BuildDFA(types.SymbolType("rare"), createAnnotatedTree("ab*c?")), args{"a"}, true},
		{"ab*c?", BuildDFA(types.SymbolType("rare"), createAnnotatedTree("ab*c?")), args{"abbb"}, true},
		{"ab*c?", BuildDFA(types.SymbolType("rare"), createAnnotatedTree("ab*c?")), args{"ac"}, true},
		{"ab*c?", BuildDFA(types.SymbolType("rare"), createAnnotatedTree("ab*c?")), args{"bbc"}, false},
		{"ab*c?(d|e)+f", BuildDFA(types.SymbolType("rare"), createAnnotatedTree("ab*c?(d|e)+f")), args{"bbc"}, false},
		{"ab*c?(d|e)+f", BuildDFA(types.SymbolType("rare"), createAnnotatedTree("ab*c?(d|e)+f")), args{"adef"}, true},
		{"ab*c?(d|e)+f", BuildDFA(types.SymbolType("rare"), createAnnotatedTree("ab*c?(d|e)+f")), args{"abbedf"}, true},
		{"ab*c?(d|e)+f", BuildDFA(types.SymbolType("rare"), createAnnotatedTree("ab*c?(d|e)+f")), args{"abbcedeededf"}, true},
		{"id1", BuildDFA(types.SymbolType("id"), createAnnotatedTree("((a|b|c|d|e|f)+|_+)((a|b|c|d|e|f)|(0|1|2|3|4))*")), args{"abbcedeededf"}, true},
		{"id2", BuildDFA(types.SymbolType("id"), createAnnotatedTree("((a|b|c|d|e|f)+|_+)((a|b|c|d|e|f)|(0|1|2|3|4))*")), args{"_____a"}, true},
		{"id3", BuildDFA(types.SymbolType("id"), createAnnotatedTree("((a|b|c|d|e|f)+|_+)((a|b|c|d|e|f)|(0|1|2|3|4))*")), args{"g"}, false},
		{"id4", BuildDFA(types.SymbolType("id"), createAnnotatedTree("((a|b|c|d|e|f)+|_+)((a|b|c|d|e|f)|(0|1|2|3|4))*")), args{"_a01"}, true},
		{"id5", BuildDFA(types.SymbolType("id"), createAnnotatedTree("((a|b|c|d|e|f)+|_+)((a|b|c|d|e|f)|(0|1|2|3|4))*")), args{"_a01g"}, false},
		{"string", BuildDFA(types.SymbolType("string"), createAnnotatedTree("\"^\"*\"")), args{"\"hello World\""}, true},
		{"string not empty", BuildDFA(types.SymbolType("string"), createAnnotatedTree("\"^\"+\"")), args{"\"\""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.dfa
			if got := d.Evaluate(tt.args.input); got != tt.want {
				t.Errorf("DFA.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
