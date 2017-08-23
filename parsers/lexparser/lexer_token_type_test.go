package lexparser

import "testing"

func Test_lexTokenType_GetPrecedence(t *testing.T) {
	tests := []struct {
		name string
		l    LexTokenType
		want int
	}{
		{"^", Carret, prefixPrecedence},
		{"⁺", Plus, postfixPrecedence},
		{"*", Asterisk, postfixPrecedence},
		{"?", Carret, prefixPrecedence},
		{"-", Minus, rangePrecedence},
		{"|", Pipe, orPrecedence},
		{"[", Leftbrace, groupbracePrecedence},
		{"(", Lefparen, groupingPrecedence},
		{"Char", Char, concatPrecedence},
		{"EOF", EOF, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.GetPrecedence(); got != tt.want {
				t.Errorf("LexTokenType.GetPrecedence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lexTokenType_isPunctuator(t *testing.T) {
	tests := []struct {
		name string
		l    LexTokenType
		want bool
	}{
		{"(", Lefparen, true},
		{")", Rightparen, true},
		{"[", Leftbrace, true},
		{"]", Rightbrace, true},
		{"|", Pipe, true},
		{"+", Plus, true},
		{"-", Minus, true},
		{"*", Asterisk, true},
		{"?", Question, true},
		{"^", Carret, true},
		{"Char", Char, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.isPunctuator(); got != tt.want {
				t.Errorf("LexTokenType.isPunctuator() = %v, want %v", got, tt.want)
			}
		})
	}
}
