package lexparser

import (
	"reflect"
	"testing"
)

func Test_rangeParselet_Parse_Error(t *testing.T) {
	getparser := func(input string, delimiters ...string) LexerParser {
		lexscanner := NewLexScanner(input)
		for _, delimiter := range delimiters {
			lexscanner.AddDelimiter(delimiter)
		}
		return NewLexerParser(lexscanner)
	}

	tests := []struct {
		name    string
		parser  LexerParser
		want    *LexerExpression
		wantErr bool
	}{
		{"expected char", getparser("[|-z]"), nil, true},
		{"expected minus", getparser("[aa-z]+"), nil, true},
		{"expected second char", getparser("[a-|]+"), nil, true},
		{"expected right brace", getparser("[a-zz]+"), nil, true},
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

func Test_prefixOperatorParselet_Parse_Error(t *testing.T) {
	getparser := func(input string, delimiters ...string) LexerParser {
		lexscanner := NewLexScanner(input)
		for _, delimiter := range delimiters {
			lexscanner.AddDelimiter(delimiter)
		}
		return NewLexerParser(lexscanner)
	}

	tests := []struct {
		name    string
		parser  LexerParser
		want    *LexerExpression
		wantErr bool
	}{
		{"expected char", getparser("^[|-z]"), nil, true},
		{"expected char", getparser("^(a|b)"), nil, true},
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

func Test_groupParselet_Parse(t *testing.T) {
	getparser := func(input string, delimiters ...string) LexerParser {
		lexscanner := NewLexScanner(input)
		for _, delimiter := range delimiters {
			lexscanner.AddDelimiter(delimiter)
		}
		return NewLexerParser(lexscanner)
	}

	tests := []struct {
		name    string
		parser  LexerParser
		want    *LexerExpression
		wantErr bool
	}{
		{"expected minus", getparser("(a+[a])"), nil, true},
		{"expected right paren", getparser("(a+*"), nil, true},
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
