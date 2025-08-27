package model

import "git.makyo.dev/makyo/gogogogogram/state"

type model struct {
	fieldSize, sectionSize, cellsPerSection int

	state *state.State

	clears, score, factor, track int

	cursor *state.Point

	columnStates, rowStates     [][]int
	columnsCorrect, rowsCorrect []bool

	history string
}

func New(sectionSize, cellsPerSection int) model {
	m := model{
		fieldSize:       sectionSize * cellsPerSection,
		sectionSize:     sectionSize,
		cellsPerSection: cellsPerSection,
		cursor:          &state.Point{0, 0},
	}
	m.state = state.New(sectionSize, cellsPerSection)

	return m
}
