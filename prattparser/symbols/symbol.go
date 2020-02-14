package symbols

// Symbol is an interface that extends the Token interface with:
// Line int
// Column int
// Value interface{}
type Symbol interface {
	Token
	Line() int
	Column() int
	Value() interface{}
}
