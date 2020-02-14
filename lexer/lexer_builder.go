package lexer

import (
	"github.com/lfmexi/golalr/lexer/dfa"
	"github.com/lfmexi/golalr/parsers/lexparser"
	"github.com/lfmexi/golalr/prattparser"
	"github.com/lfmexi/golalr/prattparser/types"
)

// Action is a type of function that is used as a callback for the Lexer's terminal definitions
// Requires a Lexer and a string
// The callback must return a ignore (bool) and an user defined value (interface{})
type Action func(*Lexer, string) (bool, interface{})

// Builder is used for creating new Lexers
type Builder struct {
	currentContext string
	annotatedTrees map[string]map[types.TokenType]*dfa.AnnotatedAST
	actions        map[string]map[types.TokenType]Action
	unknownAction  Action
	firstLine      int
	firstColumn    int
}

// NewLexerBuilder creates a new Builder of Lexer type values
func NewLexerBuilder() Builder {
	trees := make(map[string]map[types.TokenType]*dfa.AnnotatedAST)
	trees["default"] = make(map[types.TokenType]*dfa.AnnotatedAST)

	actions := make(map[string]map[types.TokenType]Action)
	actions["default"] = make(map[types.TokenType]Action)
	return Builder{
		"default",
		trees,
		actions,
		nil,
		0,
		0,
	}
}

// SetFirstLine sets the initial value for the line counter
func (b *Builder) SetFirstLine(line int) {
	b.firstLine = line
}

// SetFirstColumn sets the initial value for the char counter
func (b *Builder) SetFirstColumn(column int) {
	b.firstColumn = column
}

// AddUnknownTokenAction adds an action that will be executed when the Lexer fails on recognizing a Symbol
func (b *Builder) AddUnknownTokenAction(action Action) {
	b.unknownAction = action
}

// AddNewContext creates a new context if it not exists. All the terminal definitions declared after this call
// will be added on the context created when the Lexer is build.
func (b *Builder) AddNewContext(context string) {
	b.currentContext = context
	if b.annotatedTrees[context] == nil && b.actions[context] == nil {
		b.actions[context] = make(map[types.TokenType]Action)
		b.annotatedTrees[context] = make(map[types.TokenType]*dfa.AnnotatedAST)
	}
}

// AddPreBuiltCharDefinition adds to the current context a new terminal definition based on a PreBuiltDefinition
func (b *Builder) AddPreBuiltCharDefinition(id types.TokenType, def PreBuiltDefinition, action func(*Lexer, string) string) {
	expression := &lexparser.LexerExpression{}
	expression.SetToken(newLexSymbol(nil, lexparser.Char, string(def), nil))
	b.annotatedTrees[b.currentContext][id] = dfa.NewAnnotatedAST(expression)
}

// AddTerminalDefinition adds to the current context a new terminal definition based on a regexp. The id string is
// the id of the terminal that will be recognized by the regexp string. The action is a callback function for the
// recognition of the terminal
func (b *Builder) AddTerminalDefinition(id types.TokenType, regexp string, action Action) {
	scanner := lexparser.NewLexScanner(regexp)
	parser := lexparser.NewLexerParser(scanner)
	expression, err := parser.Parse()

	if err != nil {
		panic(err)
	}

	expression, err = expandExpression(expression)

	if err != nil {
		panic(err)
	}

	b.annotatedTrees[b.currentContext][id] = dfa.NewAnnotatedAST(expression)

	if action != nil {
		b.actions[b.currentContext][id] = action
	}
}

// Build creates a new Lexer typed value
func (b *Builder) Build(input string) prattparser.TokenIterator {
	dfas := make(map[string][]*dfa.DFA)
	actions := make(map[string]map[types.TokenType]Action)
	for context, asts := range b.annotatedTrees {
		dfas[context] = make([]*dfa.DFA, 0)
		actions[context] = make(map[types.TokenType]Action)
		for key, ast := range asts {
			dfas[context] = append(dfas[context], dfa.BuildDFA(key, ast))
			actions[context][key] = b.actions[context][key]
		}
	}

	return &Lexer{
		"default",
		0,
		input,
		dfas,
		actions,
		b.unknownAction,
		b.firstLine,
		b.firstColumn,
	}
}
