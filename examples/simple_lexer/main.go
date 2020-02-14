package main

import (
	"fmt"

	"github.com/lfmexi/golalr/lexer"
	"github.com/lfmexi/golalr/prattparser/symbols"
	"github.com/lfmexi/golalr/prattparser/types"
)

func main() {
	builder := lexer.NewLexerBuilder()

	const (
		blank        types.SymbolType = "blank"
		openTagDef   types.SymbolType = "openTagDef"
		id           types.SymbolType = "id"
		closeTagDef  types.SymbolType = "closeTagDef"
		closeTag     types.SymbolType = "closeTag"
		text         types.SymbolType = "text"
		openCloseDef types.SymbolType = "openCloseDef"
	)

	builder.AddTerminalDefinition(blank, " ", func(l *lexer.Lexer, s string) (bool, interface{}) {
		return true, nil
	})

	builder.AddTerminalDefinition(openTagDef, "<", nil)
	builder.AddTerminalDefinition(id, "([a-z]|[A-Z])([a-z]|[A-Z]|[0-9])*", nil)
	builder.AddTerminalDefinition(closeTagDef, ">", func(l *lexer.Lexer, s string) (bool, interface{}) {
		l.SetNextContext("inner")
		return false, nil
	})
	builder.AddTerminalDefinition(closeTag, "/>", nil)

	// New Inner strings context
	builder.AddNewContext("inner")
	builder.AddTerminalDefinition(text, "(^<)*", nil)
	builder.AddTerminalDefinition(openCloseDef, "</", func(l *lexer.Lexer, s string) (bool, interface{}) {
		l.SetNextContext("default")
		return false, nil
	})

	l := builder.Build("<a >Hello World</a>")
	sym := l.Next().(symbols.Symbol)
	fmt.Printf("Symbol: %v Line: %v, Column: %v, Value: %v\n", sym.TokenType(), sym.Line(), sym.Column(), sym.Text())
	sym = l.Next().(symbols.Symbol)
	fmt.Printf("Symbol: %v Line: %v, Column: %v, Value: %v\n", sym.TokenType(), sym.Line(), sym.Column(), sym.Text())
	sym = l.Next().(symbols.Symbol)
	fmt.Printf("Symbol: %v Line: %v, Column: %v, Value: %v\n", sym.TokenType(), sym.Line(), sym.Column(), sym.Text())
	sym = l.Next().(symbols.Symbol)
	fmt.Printf("Symbol: %v Line: %v, Column: %v, Value: %v\n", sym.TokenType(), sym.Line(), sym.Column(), sym.Text())
	sym = l.Next().(symbols.Symbol)
	fmt.Printf("Symbol: %v Line: %v, Column: %v, Value: %v\n", sym.TokenType(), sym.Line(), sym.Column(), sym.Text())
	sym = l.Next().(symbols.Symbol)
	fmt.Printf("Symbol: %v Line: %v, Column: %v, Value: %v\n", sym.TokenType(), sym.Line(), sym.Column(), sym.Text())
	sym = l.Next().(symbols.Symbol)
	fmt.Printf("Symbol: %v Line: %v, Column: %v, Value: %v\n", sym.TokenType(), sym.Line(), sym.Column(), sym.Text())
	sym = l.Next().(symbols.Symbol)
	fmt.Printf("Symbol: %v Line: %v, Column: %v, Value: %v\n", sym.TokenType(), sym.Line(), sym.Column(), sym.Text())
	sym = l.Next().(symbols.Symbol)
	fmt.Printf("Symbol: %v Line: %v, Column: %v, Value: %v\n", sym.TokenType(), sym.Line(), sym.Column(), sym.Text())
}
