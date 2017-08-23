package lexparser

import (
	"reflect"
	"testing"
)

func Test_charInfixParselet_Parse_Error(t *testing.T) {
	getparser := func(input string, delimiters ...string) LexerParser {
		lexscanner := NewLexScanner(input)
		for _, delimiter := range delimiters {
			lexscanner.Delimiters[delimiter] = delimiter
		}
		return NewLexerParser(lexscanner)
	}

	tests := []struct {
		name    string
		parser  LexerParser
		want    *LexerExpression
		wantErr bool
	}{
		{"expected char", getparser("ab|a-1"), nil, true},
		{"unexpected )", getparser("abc*a)"), nil, true},
		{"unexpected )", getparser("ab|()"), nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.parser.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("rangeParselet.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rangeParselet.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryOperatorParselet_Parse(t *testing.T) {
	getparser := func(input string, delimiters ...string) LexerParser {
		lexscanner := NewLexScanner(input)
		for _, delimiter := range delimiters {
			lexscanner.Delimiters[delimiter] = delimiter
		}
		return NewLexerParser(lexscanner)
	}

	tests := []struct {
		name    string
		parser  LexerParser
		want    *LexerExpression
		wantErr bool
	}{
		{"expected char", getparser("a||"), nil, true},
		{"unexpected eof", getparser("a|"), nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.parser.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("rangeParselet.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rangeParselet.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rightComplexInfixParselet_Parse(t *testing.T) {
	getparser := func(input string, delimiters ...string) LexerParser {
		lexscanner := NewLexScanner(input)
		for _, delimiter := range delimiters {
			lexscanner.Delimiters[delimiter] = delimiter
		}
		return NewLexerParser(lexscanner)
	}

	tests := []struct {
		name    string
		parser  LexerParser
		want    *LexerExpression
		wantErr bool
	}{
		{"unexpected char", getparser("a|[aa-z]"), nil, true},
		{"unexpected char after minus", getparser("a|[a-zz]"), nil, true},
		{"unexpected asterisk", getparser("a|[a*a"), nil, true},
		{"error on group", getparser("a(b[az])"), nil, true},
		{"error after group", getparser("a(b[a-z])^"), nil, true},
		{"error in group", getparser("a(b(a|a)"), nil, true},
		{"error on next expression", getparser("a(b^[a-z])[az]"), nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.parser.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("rangeParselet.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rangeParselet.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
