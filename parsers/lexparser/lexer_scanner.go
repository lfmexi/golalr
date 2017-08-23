package lexparser

import "github.com/lfmexi/golalr/symbols"

// LexScanner is the struct that implements the prattparser.TokenInterface iterator.
// Exposes the Delimiters map, which contains the characters that should be ignored as
// delimiters
type LexScanner struct {
	index       int
	text        string
	punctuators map[rune]LexTokenType
	Delimiters  map[string]string
}

// NewLexScanner returns a pointer to LexScanner based on an input string
func NewLexScanner(text string) *LexScanner {
	scanner := &LexScanner{
		0,
		text,
		make(map[rune]LexTokenType),
		make(map[string]string),
	}
	scanner.fillPunctuators()
	return scanner
}

func (l *LexScanner) fillPunctuators() {
	l.punctuators['('] = Lefparen
	l.punctuators[')'] = Rightparen
	l.punctuators['['] = Leftbrace
	l.punctuators[']'] = Rightbrace
	l.punctuators['|'] = Pipe
	l.punctuators['-'] = Minus
	l.punctuators['+'] = Plus
	l.punctuators['*'] = Asterisk
	l.punctuators['?'] = Question
	l.punctuators['^'] = Carret
	l.punctuators['·'] = Middledot
}

// Next obtains the reference to the next symbols.Token on the given input string
func (l *LexScanner) Next() symbols.Token {
	bytes := []byte(l.text)

	for l.index < len(bytes) {
		character := bytes[l.index]
		l.index++
		punctuatorChar := rune(character)
		punct := l.punctuators[punctuatorChar]
		if punct != 0 && punct.isPunctuator() {
			return &LexToken{string(punctuatorChar), punct}
		} else if delim := l.Delimiters[string(character)]; delim == "" {
			return &LexToken{string(character), Char}
		}
	}

	return &LexToken{"EOF", EOF}
}
