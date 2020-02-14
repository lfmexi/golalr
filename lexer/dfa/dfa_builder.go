package dfa

import "github.com/lfmexi/golalr/prattparser/types"

// BuildDFA creates a DFA using an AnnotatedAST for the creation of
// its inner states
func BuildDFA(terminal types.TokenType, annotatedTree *AnnotatedAST) *DFA {
	annotatedTree.root.setFirstPositions()
	annotatedTree.root.setLastPositions()
	annotatedTree.root.setNextPositions(make([]int, 0), annotatedTree.root.lastPositions)

	initialState, states := createStates(annotatedTree)
	return &DFA{
		terminal,
		initialState,
		states,
	}
}

func createStates(annotatedTree *AnnotatedAST) (string, map[string]*state) {
	states := make(map[string]*state)
	st := newState(annotatedTree.root.firstPositions)
	states[st.label] = st
	initialState := st.label
	for st != nil {
		st.marked = true
		for symbol, pos := range annotatedTree.symbols {
			positions := getPositionsInState(st.positions, pos)
			sliceOfNext := getSliceOfNextFromLeafs(annotatedTree.leafs, positions)
			nextState := mergeSliceOfNext(sliceOfNext)
			nstate := newState(nextState)
			if len(nextState) > 0 {
				if states[nstate.label] == nil {
					states[nstate.label] = nstate
				}
				st.transitions[symbol] = states[nstate.label]
			}
			if symbol == "EOF" && len(positions) > 0 {
				st.acceptance = true
			}
		}
		st = getUnmarkedState(states)
	}
	return initialState, states
}

func getUnmarkedState(states map[string]*state) *state {
	for _, st := range states {
		if !st.marked {
			return st
		}
	}
	return nil
}

func mergeSliceOfNext(next [][]int) []int {
	newNext := make([]int, 0)
	for _, slice := range next {
		for _, pos := range slice {
			if !contains(newNext, pos) {
				newNext = append(newNext, pos)
			}
		}
	}
	return newNext
}

func getSliceOfNextFromLeafs(leafs map[int]*annotatedASTNode, positions []int) [][]int {
	nextSlice := make([][]int, 0)
	for _, p := range positions {
		nextSlice = append(nextSlice, leafs[p].nextPositions)
	}
	return nextSlice
}

func getPositionsInState(statePositions []int, symbolPositions map[int]int) []int {
	positions := make([]int, 0)
	for _, p := range statePositions {
		if pos := symbolPositions[p]; pos != 0 {
			positions = append(positions, pos)
		}
	}
	return positions
}
