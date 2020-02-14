package lexer

import (
	"fmt"

	"github.com/lfmexi/golalr/lexer/dfa"
	"github.com/lfmexi/golalr/prattparser/symbols"
	"github.com/lfmexi/golalr/prattparser/types"
)

// Lexer is an implementation of TokenIterator
// It manages inner collections of state machines for Token regonition
// Has an inner context wich indicates on which collection of state machines the input will be scanned
type Lexer struct {
	context       string
	index         int
	input         string
	dfas          map[string][]*dfa.DFA
	actions       map[string]map[types.TokenType]Action
	unknownAction Action
	line          int
	column        int
}

// SetNextContext allows to change the current context at runtime
func (l *Lexer) SetNextContext(context string) {
	if l.dfas[context] == nil {
		panic(fmt.Errorf("Context %s is not defined", context))
	}
	l.context = context
}

// ResetLinesTo set the lines counter to the provided int
func (l *Lexer) ResetLinesTo(start int) {
	l.line = start
}

// IncrementLine increments the line counter by 1
func (l *Lexer) IncrementLine() {
	l.line++
}

// ResetColumnsTo set the char counter to the provided int
func (l *Lexer) ResetColumnsTo(start int) {
	l.column = start
}

// Next returns the next symbol.Symbol recognized by the Lexer
// If the input is empty it will return a types.EOF typed Symbol
// If the Lexer cannot recognize the input, it will return a types.UnknownSymbol typed Symbol
func (l *Lexer) Next() symbols.Token {
	start := l.index

	if l.index >= len(l.input) {
		return newLexSymbol(l, types.EOF, "", nil)
	}

	var substr string
	for l.index < len(l.input) {
		substr = l.input[start : l.index+1]
		for i := 0; i < len(l.dfas[l.context]); i++ {
			d := l.dfas[l.context][i]
			if d.Evaluate(substr) {
				l.index++
				terminal := l.dfas[l.context][i].GetTerminal()
				var ignore bool
				var symbolValue interface{}

				if l.index >= len(l.input) {
					if l.actions[l.context][terminal] != nil {
						ignore, symbolValue = l.actions[l.context][terminal](l, substr)
					}
					return newLexSymbol(l, terminal, substr, symbolValue)
				}

				newSubs := l.input[start : l.index+1]
				if !l.preEvaluate(newSubs) {
					if l.actions[l.context][terminal] != nil {
						ignore, symbolValue = l.actions[l.context][terminal](l, substr)
					}
					if ignore {
						l.column = l.index - start + l.column
						return l.Next()
					}
					symbol := newLexSymbol(l, terminal, substr, symbolValue)

					l.column = l.index - start + l.column
					return symbol
				}

				l.index--
				break
			}
		}
		l.index++
	}

	var unknownActionResult interface{}
	if l.unknownAction != nil {
		_, unknownActionResult = l.unknownAction(l, substr)
	}
	return newLexSymbol(l, types.UnknownSymbol, substr, unknownActionResult)
}

// AddDelimiter not implemented
func (l *Lexer) AddDelimiter(del string) {
	// Not implemented
}

func (l Lexer) preEvaluate(str string) bool {
	dfas := l.dfas[l.context]
	for _, d := range dfas {
		if d.Evaluate(str) {
			return true
		}
	}
	return false
}
