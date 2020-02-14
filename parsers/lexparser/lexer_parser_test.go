package lexparser

import "testing"
import "reflect"

func TestNewLexerParser(t *testing.T) {
	lexerscanner := NewLexScanner("ab|c")
	lexparser := NewLexerParser(lexerscanner)
	parser := lexparser.GetInnerParser()
	if parser.GetPrefixParselet(Char) == nil {
		t.Errorf("Expected to have a prefix parselet for Char")
	}

	if parser.GetPrefixParselet(Lefparen) == nil {
		t.Errorf("Expected to have a prefix parselet for Lefparen")
	}

	if parser.GetPrefixParselet(Leftbrace) == nil {
		t.Errorf("Expected to have a prefix parselet for Leftbrace")
	}

	if parser.GetPrefixParselet(Carret) == nil {
		t.Errorf("Expected to have a prefix parselet for Carret")
	}

	if parser.GetInfixParselet(Plus) == nil {
		t.Errorf("Expected to have an infix parselet for Plus")
	}

	if parser.GetInfixParselet(Asterisk) == nil {
		t.Errorf("Expected to have an infix parselet for Asterisk")
	}

	if parser.GetInfixParselet(Question) == nil {
		t.Errorf("Expected to have an infix parselet for Question")
	}

	if parser.GetInfixParselet(Pipe) == nil {
		t.Errorf("Expected to have an infix parselet for Pipe")
	}

	if parser.GetInfixParselet(Char) == nil {
		t.Errorf("Expected to have an infix parselet for Char")
	}

	if parser.GetInfixParselet(Leftbrace) == nil {
		t.Errorf("Expected to have an infix parselet for Leftbrace")
	}

	if parser.GetInfixParselet(Lefparen) == nil {
		t.Errorf("Expected to have an infix parselet for Lefparen")
	}
}

func TestLexerParser_Parse(t *testing.T) {
	getparser := func(input string, delimiters ...string) LexerParser {
		lexscanner := NewLexScanner(input)
		for _, delimiter := range delimiters {
			lexscanner.AddDelimiter(delimiter)
		}
		return NewLexerParser(lexscanner)
	}

	firstTree := "{\"${concat}\": {\"left\":\"a\",\"right\":{\"+\": {\"left\":\"b\"}}}}"

	composedTree := "{\"|\": {\"left\":{\"${concat}\": {\"left\":\"a\",\"right\":\"b\"}},\"right\":{\"*\": {\"left\":\"c\"}}}}"

	numberTree := "{\"+\": {\"left\":{\"-\": {\"left\":\"0\",\"right\":\"9\"}}}}"

	idTree := "{\"${concat}\": {\"left\":{\"|\": {\"left\":{\"-\": {\"left\":\"a\",\"right\":\"z\"}},\"right\":{\"-\": {\"left\":\"A\",\"right\":\"Z\"}}}},\"right\":{\"*\": {\"left\":{\"|\": {\"left\":{\"-\": {\"left\":\"a\",\"right\":\"z\"}},\"right\":{\"-\": {\"left\":\"A\",\"right\":\"Z\"}}}}}}}}"

	multiGroup := "{\"${concat}\": {\"left\":{\"${concat}\": {\"left\":\"a\",\"right\":{\"*\": {\"left\":\"z\"}}}},\"right\":{\"|\": {\"left\":\"a\",\"right\":\"z\"}}}}"

	stringCase := "{\"${concat}\": {\"left\":{\"${concat}\": {\"left\":\"\"\",\"right\":{\"*\": {\"left\":{\"^\": {\"right\":\"\"\"}}}}}},\"right\":\"\"\"}}"

	tests := []struct {
		name   string
		parser LexerParser
		want   string
	}{
		{"simpleRegexp", getparser("ab+", " ", "\n"), firstTree},
		{"composedRegexp", getparser("(ab)|c*", " ", "\n"), composedTree},
		{"numberRegexp", getparser("[0-9]+", " ", "\n"), numberTree},
		{"idRegexp", getparser("[a-z]|[A-Z]([a-z]|[A-Z])*"), idTree},
		{"multigroup", getparser("(az*)(a|z)"), multiGroup},
		{"stringcase", getparser("\"(^\")*\""), stringCase},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exp, err := tt.parser.Parse()
			if err != nil {
				t.Error(err)
				t.Fail()
				return
			}
			if got := exp.ToString(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LexerParser.Parse().ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
