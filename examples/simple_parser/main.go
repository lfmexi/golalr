package main

import "github.com/lfmexi/golalr/lexer"
import "github.com/lfmexi/golalr/parsers/grammarparser"
import "fmt"

func main() {
	grammarLexerBuilder := lexer.NewLexerBuilder()
	grammarLexerBuilder.AddTerminalDefinition(grammarparser.SymbolID, "([a-z]|[A-Z])([a-z]|[A-Z]|[0-9])*", nil)

	var blank grammarparser.ProductionTokenType = "blank"
	grammarLexerBuilder.AddTerminalDefinition(blank, " ", func(l *lexer.Lexer, s string) (bool, interface{}) {
		return true, nil
	})

	input := "E plus T"
	scanner := grammarLexerBuilder.Build(input)
	parser := grammarparser.NewProductionParser(&scanner)
	expression, _ := parser.Parse()
	fmt.Println(expression.Token().Text())
	expression = expression.Next()
	if expression != nil {
		fmt.Println(expression.Token().Text())
	}
	expression = expression.Next()
	if expression != nil {
		fmt.Println(expression.Token().Text())
	}

	expression = expression.Next()
	if expression == nil {
		fmt.Println("EOF my friend")
	}
}
