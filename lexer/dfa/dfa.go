package dfa

import "github.com/lfmexi/golalr/types"

// DFA represents a Deterministic Finite Automaton
type DFA struct {
	terminal     types.TokenType
	initialState string
	states       map[string]*state
}

// GetTerminal returns the id of the terminal
// represented by this DFA
func (d DFA) GetTerminal() types.TokenType {
	return d.terminal
}

// Evaluate tells if an input string is accepted
// by this DFA
func (d DFA) Evaluate(input string) bool {
	st := d.states[d.initialState]
	bytes := []byte(input)
	for i := 0; i < len(bytes); i++ {
		char := string(bytes[i])

		anyChar := st.isAnyChar(char[0])
		if st.transitions[char] == nil && anyChar == "" {
			break
		}
		if anyChar != "" {
			char = anyChar
		}
		st = st.transitions[char]
		if i == len(bytes)-1 {
			return st.acceptance
		}
	}
	return false
}
