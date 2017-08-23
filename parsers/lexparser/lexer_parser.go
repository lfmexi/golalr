package lexparser

import "github.com/lfmexi/golalr/prattparser"

// LexerParser is an extension of parattparser.Parser
// Register specific parselets for each of the operators contained
// in a regular expression.
type LexerParser struct {
	innerParser *prattparser.Parser
}

// NewLexerParser creates a new LexerParser, register its parselets
// and returns it
func NewLexerParser(l *LexScanner) LexerParser {
	parser := prattparser.NewParser(l)

	parser.RegisterPrefixParselet(Char, &simpleParselet{})

	parser.RegisterPrefixParselet(Lefparen, &groupParselet{})
	parser.RegisterPrefixParselet(Leftbrace, &rangeParselet{})

	parser.RegisterPrefixParselet(Carret, &prefixOperatorParselet{})

	parser.RegisterInfixParselet(Plus, &postfixOperatorParselet{})
	parser.RegisterInfixParselet(Asterisk, &postfixOperatorParselet{})
	parser.RegisterInfixParselet(Question, &postfixOperatorParselet{})

	parser.RegisterInfixParselet(Pipe, &binaryOperatorParselet{})

	parser.RegisterInfixParselet(Char, &charInfixParselet{})
	parser.RegisterInfixParselet(Lefparen, &rightComplexInfixParselet{})
	parser.RegisterInfixParselet(Leftbrace, &rightComplexInfixParselet{})
	parser.RegisterInfixParselet(Carret, &rightComplexInfixParselet{})

	return LexerParser{&parser}
}

// Parse generates a new LexerExpression by parsing the stream of tokens given
// by a inner prattparser.Parser
func (l *LexerParser) Parse() (*LexerExpression, error) {
	expression, err := l.innerParser.Parse(0)

	if err != nil {
		return nil, err
	}

	lexexpression := expression.(LexerExpression)
	return &lexexpression, nil
}

// GetInnerParser exposes the inner parser
func (l *LexerParser) GetInnerParser() *prattparser.Parser {
	return l.innerParser
}
