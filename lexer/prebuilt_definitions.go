package lexer

// PreBuiltDefinition is a string based type
// It represents all the pre built definitions that can be used as part of a Language
type PreBuiltDefinition string

// LeftParen
// RightParen
// LeftSquareBracket
// RightSqureBracket
// LeftCurlyBrace
// RightCurlyBrace
// Minus
// Asterisk
// Slash
// Plus
// Carret
// QuestionMark
// Bang
const (
	LeftParen         PreBuiltDefinition = "("
	RightParen        PreBuiltDefinition = ")"
	LeftSquareBracket PreBuiltDefinition = "["
	RightSqureBracket PreBuiltDefinition = "]"
	LeftCurlyBrace    PreBuiltDefinition = "{"
	RightCurlyBrace   PreBuiltDefinition = "}"
	Minus             PreBuiltDefinition = "-"
	Asterisk          PreBuiltDefinition = "*"
	Slash             PreBuiltDefinition = "/"
	Plus              PreBuiltDefinition = "+"
	Carret            PreBuiltDefinition = "^"
	QuestionMark      PreBuiltDefinition = "?"
	Bang              PreBuiltDefinition = "!"
)
