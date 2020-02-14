package lexer

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lfmexi/golalr/prattparser"

	"github.com/lfmexi/golalr/prattparser/symbols"
	"github.com/lfmexi/golalr/prattparser/types"
)

const (
	plus       types.SymbolType = "plus"
	minus      types.SymbolType = "minus"
	div        types.SymbolType = "div"
	factor     types.SymbolType = "factor"
	leftParen  types.SymbolType = "leftParen"
	rightParen types.SymbolType = "rightParen"
	space      types.SymbolType = "space"
	number     types.SymbolType = "number"

	openTagDef   types.SymbolType = "openTagDef"
	id           types.SymbolType = "id"
	closeTagDef  types.SymbolType = "closeTagDef"
	closeTag     types.SymbolType = "closeTag"
	text         types.SymbolType = "text"
	openCloseDef types.SymbolType = "openCloseDef"
	newLine      types.SymbolType = "newLine"
	EOF          types.SymbolType = "EOF"
)

// Recognizes the language members of arithmetic operations
func newArithmeticBuilder() Builder {
	builder := NewLexerBuilder()
	builder.AddPreBuiltCharDefinition(plus, Plus, nil)
	builder.AddPreBuiltCharDefinition(minus, Minus, nil)
	builder.AddPreBuiltCharDefinition(div, Slash, nil)
	builder.AddPreBuiltCharDefinition(factor, Asterisk, nil)
	builder.AddPreBuiltCharDefinition(leftParen, LeftParen, nil)
	builder.AddPreBuiltCharDefinition(rightParen, RightParen, nil)

	builder.AddUnknownTokenAction(func(l *Lexer, s string) (bool, interface{}) {
		return false, fmt.Errorf("Lexical error: %v at %v", s, l.column)
	})

	// This will ignore the spaces
	builder.AddTerminalDefinition(space, " ", func(*Lexer, string) (bool, interface{}) {
		return true, nil
	})

	builder.AddTerminalDefinition(number, "[0-9]+(.[0-9]*)?", nil)
	return builder
}

func newMarkupLexerBuilder() Builder {
	builder := NewLexerBuilder()
	ignoreNewLine := func(l *Lexer, s string) (bool, interface{}) {
		l.IncrementLine()
		l.ResetColumnsTo(1)
		return true, nil
	}
	builder.SetFirstColumn(1)
	builder.SetFirstLine(1)
	builder.AddTerminalDefinition(openTagDef, "<", nil)
	builder.AddTerminalDefinition(id, "([a-z]|[A-Z])([a-z]|[A-Z]|[0-9])*", nil)
	builder.AddTerminalDefinition(closeTagDef, ">", func(l *Lexer, s string) (bool, interface{}) {
		l.SetNextContext("inner")
		return false, nil
	})
	builder.AddTerminalDefinition(closeTag, "/>", nil)
	builder.AddTerminalDefinition(newLine, "\n", ignoreNewLine)

	// New Inner strings context
	builder.AddNewContext("inner")
	builder.AddTerminalDefinition(text, "(^<)*", nil)
	builder.AddTerminalDefinition(openCloseDef, "</", func(l *Lexer, s string) (bool, interface{}) {
		l.SetNextContext("default")
		l.ResetLinesTo(1)
		return false, nil
	})

	return builder
}

func newWrongLexerBuilder() Builder {
	builder := NewLexerBuilder()
	const a types.SymbolType = "a"
	builder.AddTerminalDefinition(a, "a*", func(l *Lexer, s string) (bool, interface{}) {
		l.SetNextContext("foo")
		return false, nil
	})
	return builder
}

func TestArithmeticLexer_Next(t *testing.T) {
	builder := newArithmeticBuilder()
	lexer := builder.Build("2+19.1*(3-4)")
	lexerForSpaced := builder.Build("      1 + 2")
	lexerForWrongInput := builder.Build("4 % 3")
	tests := []struct {
		name  string
		lexer prattparser.TokenIterator
		want  string
	}{
		{"Number", lexer, "number"},
		{"Plus", lexer, "plus"},
		{"Number", lexer, "number"},
		{"Factor", lexer, "factor"},
		{"LeftParen", lexer, "leftParen"},
		{"Number", lexer, "number"},
		{"Minus", lexer, "minus"},
		{"Number", lexer, "number"},
		{"RightParen", lexer, "rightParen"},
		{"EOF", lexer, "EOF"},
		{"number", lexerForSpaced, "number"},
		{"number", lexerForWrongInput, "number"},
		{"wrongInput", lexerForWrongInput, "UNKNOWN_SYMBOL"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lexer.Next(); !reflect.DeepEqual(got.TokenType(), types.SymbolType(tt.want)) {
				t.Errorf("Lexer.Next() = %v, want %v", got.TokenType(), types.SymbolType(tt.want))
			}
		})
	}
}

func TestMultiContextLexer_Next(t *testing.T) {
	builder := newMarkupLexerBuilder()
	lexer := builder.Build("<a\nb>Hello world,This is a new line</a>")
	tests := []struct {
		name   string
		lexer  prattparser.TokenIterator
		want   types.TokenType
		column int
		line   int
		value  interface{}
	}{
		{string(openTagDef), lexer, openTagDef, 1, 1, nil},
		{string(id), lexer, id, 2, 1, nil},
		{string(id), lexer, id, 2, 2, nil},
		{string(closeTagDef), lexer, closeTagDef, 3, 2, nil},
		{string(text), lexer, text, 4, 2, nil},
		{string(openCloseDef), lexer, openCloseDef, 34, 1, nil},
		{string(id), lexer, id, 36, 1, nil},
		{string(closeTagDef), lexer, closeTagDef, 37, 1, nil},
		{string(EOF), lexer, EOF, 37, 1, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lexer.Next()
			if !reflect.DeepEqual(got.TokenType(), tt.want) {
				t.Errorf("Lexer.Next() = %v, want %v", got.TokenType(), tt.want)
			}
			if got.(symbols.Symbol).Column() != tt.column {
				t.Errorf("Lexer.Next().Column = %v, want %v", got.(symbols.Symbol).Column(), tt.column)
			}
			if got.(symbols.Symbol).Line() != tt.line {
				t.Errorf("Lexer.Next().Line = %v, want %v", got.(symbols.Symbol).Line(), tt.line)
			}
			if !reflect.DeepEqual(got.(symbols.Symbol).Value(), tt.value) {
				t.Errorf("Lexer.Next().Value = %v, want %v", got.(symbols.Symbol).Value(), tt.value)
			}
		})
	}
}

func TestWrongLexer_Next(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Did not recover")
		}
	}()
	builder := newWrongLexerBuilder()
	lexer := builder.Build("aaa")
	lexer.Next()
}
