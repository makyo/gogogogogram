package model

type point struct {
	x, y int
}

type model struct {
	fieldSize, sectionSize, cellsPerSection int

	state *state

	clears, score, factor, track int

	cursor *point

	columnStates, rowStates     [][]int
	columnsCorrect, rowsCorrect []bool

	history string
}

func New(sectionSize, cellsPerSection int) model {
	m := model{
		fieldSize:       sectionSize * cellsPerSection,
		sectionSize:     sectionSize,
		cellsPerSection: cellsPerSection,
		cursor:          &point{0, 0},
	}
	m.state = newState(sectionSize, cellsPerSection)

	return m
}
