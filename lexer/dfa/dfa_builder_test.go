package dfa

import (
	"testing"

	"github.com/lfmexi/golalr/parsers/lexparser"
	"github.com/lfmexi/golalr/prattparser/types"
)

func TestBuildDFA_WithSuccessfulValidations(t *testing.T) {
	scanner := lexparser.NewLexScanner("(a|b)*abb")
	parser := lexparser.NewLexerParser(scanner)
	expression, _ := parser.Parse()
	annotatedAST := NewAnnotatedAST(expression)

	scanner = lexparser.NewLexScanner("0|1|2|3")
	parser = lexparser.NewLexerParser(scanner)
	expression, _ = parser.Parse()
	secondAnnotatedAST := NewAnnotatedAST(expression)

	scanner = lexparser.NewLexScanner("(0|1|2|3|4|5|6|7|8|9)+")
	parser = lexparser.NewLexerParser(scanner)
	expression, _ = parser.Parse()
	thirdAnnotatedAST := NewAnnotatedAST(expression)

	type args struct {
		terminal      types.TokenType
		annotatedTree *AnnotatedAST
	}

	tests := []struct {
		name string
		args args
	}{
		{"For (a|b)*abb", args{types.SymbolType("sample"), annotatedAST}},
		{"For 0|1|2|3", args{types.SymbolType("digit"), secondAnnotatedAST}},
		{"For (0|1|2|3|4|5|6|7|8|9)+", args{types.SymbolType("integer"), thirdAnnotatedAST}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildDFA(tt.args.terminal, tt.args.annotatedTree)
			if got == nil {
				t.Errorf("BuildDFA() != %v", got)
			}
		})
	}
}
