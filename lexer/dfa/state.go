package dfa

import (
	"fmt"
	"strconv"
)

type state struct {
	label       string
	marked      bool
	acceptance  bool
	positions   []int
	transitions map[string]*state
}

func newState(positions []int) *state {
	var label string
	for i := range positions {
		label = fmt.Sprintf("%v%v", label, strconv.Itoa(positions[i]))
	}
	return &state{
		label,
		false,
		false,
		positions,
		make(map[string]*state),
	}
}

func (st state) isAnyChar(char byte) string {
	for transition := range st.transitions {
		if transition[0] == '^' && len(transition) > 1 {
			if transition[1] != char {
				return transition
			}
		}
	}
	return ""
}
