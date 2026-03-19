package ui

import (
	"os"

	"git.makyo.dev/makyo/gogogogogram/state"
)

type model struct {
	fieldSize, sectionSize, cellsPerSection int

	state *state.State

	filename string
	file     *os.File

	clears, score, factor, track int

	columnStates, rowStates     [][]int
	columnsCorrect, rowsCorrect []bool
}

func NewModel(sectionSize, cellsPerSection int) model {
	m := model{
		fieldSize:       sectionSize * cellsPerSection,
		sectionSize:     sectionSize,
		cellsPerSection: cellsPerSection,
		state:           state.New(sectionSize, cellsPerSection),
	}

	return m
}
