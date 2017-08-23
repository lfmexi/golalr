package prattparser

import (
	"fmt"
	"testing"

	"github.com/lfmexi/golalr/symbols"
	"github.com/lfmexi/golalr/types"
)

type tokenType int

func (t tokenType) GetPrecedence() int {
	return int(t)
}

const foo tokenType = 12

type simpleASTNode struct{}

func (simpleASTNode) IsLeaf() bool {
	return true
}

type simpleExpression struct{}

type simplePrefixParselet struct{}

func (simplePrefixParselet) Parse(p *Parser, t symbols.Token) (Expression, error) {
	return &simpleExpression{}, nil
}

type simpleInfixParselet struct{}

func (s simpleInfixParselet) Parse(p *Parser, left Expression, t symbols.Token) (Expression, error) {
	return p.Parse(t.TokenType().GetPrecedence())
}

type testToken struct {
	text  string
	tType types.TokenType
}

func (t testToken) TokenType() types.TokenType {
	return t.tType
}

func (t testToken) Text() string {
	return t.text
}

type tokenIterator struct{}

func (*tokenIterator) Next() symbols.Token {
	return &testToken{text: "foo", tType: foo}
}

func TestNewParser(t *testing.T) {
	fmt.Println("Test NewParser")
	parser := NewParser(&tokenIterator{})
	if &parser == nil {
		t.Fail()
	}
}

func TestRegisterPrefixParselet(t *testing.T) {
	fmt.Println("Test RegisterPrefixParselet")
	parser := NewParser(&tokenIterator{})
	var tType tokenType
	parser.RegisterPrefixParselet(tType, &simplePrefixParselet{})

	if parser.prefixParselets[tType] == nil {
		t.Fail()
	}
}

func TestRegisterInfixParselet(t *testing.T) {
	fmt.Println("Test RegisterInfixParselet")
	parser := NewParser(&tokenIterator{})
	var tType tokenType = 1
	parser.RegisterInfixParselet(tType, &simpleInfixParselet{})

	if parser.infixParselets[tType] == nil {
		t.Fail()
	}
}

func TestLookNext(t *testing.T) {
	fmt.Println("Test LookNext")
	parser := NewParser(&tokenIterator{})
	token := parser.LookNext()
	if token == nil {
		t.Fail()
	}
	if token.TokenType() != foo {
		t.Fail()
	}
}

func TestBufferingByLookingNext(t *testing.T) {
	fmt.Println("Test Buffering by LookNext")
	parser := NewParser(&tokenIterator{})
	parser.LookNext()
	parser.LookNext()
	parser.LookNext()
	parser.LookNext()
	if len(parser.tokensBuffer) != 1 {
		t.Fail()
	}
}

const (
	cero  tokenType = 0
	one   tokenType = 1
	two   tokenType = 2
	three tokenType = 3
	EOF   tokenType = 0
)

type customIterator struct {
	tokens []tokenType
}

func (c *customIterator) Next() symbols.Token {
	fmt.Println(c.tokens)
	if len(c.tokens) > 0 {
		token := &testToken{"foo", c.tokens[0]}
		c.tokens = append(c.tokens[:0], c.tokens[1:]...)
		return token
	}
	return &testToken{"", EOF}
}

func TestConsume(t *testing.T) {
	fmt.Println("Test Consume")
	expectedTypes := []tokenType{one, two, three}
	parser := NewParser(&customIterator{expectedTypes})
	token1, _ := parser.Consume(nil)
	if token1 == nil && token1.TokenType() != one {
		t.Fail()
	}
	token2, _ := parser.Consume(nil)
	if token2 == nil && token2.TokenType() != two {
		t.Fail()
	}
	token3, _ := parser.Consume(nil)
	if token3 == nil && token3.TokenType() != three {
		t.Fail()
	}
	token4, _ := parser.Consume(nil)
	if token4 == nil && token4.TokenType() != EOF {
		t.Fail()
	}
}

func TestConsumeType(t *testing.T) {
	fmt.Println("Test Consumen by Type")
	expectedTypes := []tokenType{foo}
	parser := NewParser(&customIterator{expectedTypes})
	if token, err := parser.Consume(foo); token == nil && err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestConsumeError(t *testing.T) {
	fmt.Println("Test Error on Consume")
	expectedTypes := []tokenType{foo}
	parser := NewParser(&customIterator{expectedTypes})
	if _, err := parser.Consume(one); err == nil {
		t.Fail()
	}
}

func TestGetPrefixParselet(t *testing.T) {
	expectedTypes := []tokenType{cero, one, cero, two}
	iterator := &customIterator{expectedTypes}
	parser := NewParser(iterator)
	simpleParselet := &simplePrefixParselet{}
	parser.RegisterPrefixParselet(cero, simpleParselet)
	if simpleParselet != parser.GetPrefixParselet(cero) {
		t.Error("GetPrefixParselet does not match with the actual parselet")
		t.Fail()
	}
}

func TestGetInfixParselet(t *testing.T) {
	expectedTypes := []tokenType{cero, one, cero, two}
	iterator := &customIterator{expectedTypes}
	parser := NewParser(iterator)
	simpleParselet := &simpleInfixParselet{}
	parser.RegisterInfixParselet(cero, simpleParselet)
	if simpleParselet != parser.GetInfixParselet(cero) {
		t.Error("GetPrefixParselet does not match with the actual parselet")
		t.Fail()
	}
}

func TestParse(t *testing.T) {
	fmt.Println("Test Parse")
	expectedTypes := []tokenType{cero, one, cero, two}
	iterator := &customIterator{expectedTypes}
	parser := NewParser(iterator)
	parser.RegisterPrefixParselet(cero, &simplePrefixParselet{})
	parser.RegisterPrefixParselet(EOF, &simplePrefixParselet{})
	parser.RegisterInfixParselet(one, &simpleInfixParselet{})
	parser.RegisterInfixParselet(two, &simpleInfixParselet{})
	if expression, err := parser.Parse(0); expression == nil && err != nil {
		t.Error(err)
		t.Fail()
	}

	if len(iterator.tokens) > 0 {
		t.Errorf("Didn't consume the stack")
		t.Fail()
	}
}

func TestParseWithErrorAtFirsPosition(t *testing.T) {
	fmt.Println("Test Parse error at first position")
	expectedTypes := []tokenType{cero, one, two}
	iterator := &customIterator{expectedTypes}
	parser := NewParser(iterator)
	parser.RegisterInfixParselet(one, &simpleInfixParselet{})
	parser.RegisterInfixParselet(two, &simpleInfixParselet{})
	if expression, err := parser.Parse(0); expression != nil && err == nil {
		t.Error(err)
		t.Fail()
	}
}

func TestParseWithErrorOnInfix(t *testing.T) {
	fmt.Println("Test Parse")
	expectedTypes := []tokenType{cero, foo, one, cero, two}
	iterator := &customIterator{expectedTypes}
	parser := NewParser(iterator)
	parser.RegisterPrefixParselet(cero, &simplePrefixParselet{})
	parser.RegisterPrefixParselet(EOF, &simplePrefixParselet{})
	parser.RegisterInfixParselet(one, &simpleInfixParselet{})
	parser.RegisterInfixParselet(two, &simpleInfixParselet{})
	expression, err := parser.Parse(0)
	if expression != nil && err == nil {
		t.Error("Parse error expected")
		t.Fail()
	}
	if err.Error() == "" {
		t.Error("Empty error")
	}
}

type infixErrorParselet struct{}

func (infixErrorParselet) Parse(o *Parser, e Expression, t symbols.Token) (Expression, error) {
	return nil, fmt.Errorf("Not recognized token %s", t.Text())
}

type prefixErrorParselet struct{}

func (prefixErrorParselet) Parse(o *Parser, t symbols.Token) (Expression, error) {
	return nil, fmt.Errorf("Not recognized token %s", t.Text())
}

func TestParseWithErrorOnPrefix(t *testing.T) {
	fmt.Println("Test Parse")
	expectedTypes := []tokenType{cero, foo, one, cero, two}
	iterator := &customIterator{expectedTypes}
	parser := NewParser(iterator)
	parser.RegisterPrefixParselet(cero, &prefixErrorParselet{})
	parser.RegisterInfixParselet(one, &simpleInfixParselet{})
	parser.RegisterInfixParselet(two, &simpleInfixParselet{})
	if expression, err := parser.Parse(0); expression != nil && err == nil {
		t.Error("Parse error expected")
		t.Fail()
	}
}

func TestParseWithErrorOnInfixParse(t *testing.T) {
	fmt.Println("Test Parse")
	expectedTypes := []tokenType{cero, two}
	iterator := &customIterator{expectedTypes}
	parser := NewParser(iterator)
	parser.RegisterPrefixParselet(cero, &simplePrefixParselet{})
	parser.RegisterPrefixParselet(EOF, &simplePrefixParselet{})
	parser.RegisterInfixParselet(one, &infixErrorParselet{})
	parser.RegisterInfixParselet(two, &infixErrorParselet{})
	if expression, err := parser.Parse(0); expression != nil && err == nil {
		t.Error("Parse error expected")
		t.Fail()
	}
}
