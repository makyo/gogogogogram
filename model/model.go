package model

type Model struct {
	size, section, perSection    int
	field, view                  int
	marks, flags                 int
	correct, sections, completed int
	clears, score, factor, track int
	cursor                       int

	columnStates, rowStates     [][]int
	columnsCorrect, rowsCorrect []bool

	history []string
}

func New(section, perSection int) (Model, error) {
	m := Model{
		size:       section * perSection,
		section:    section,
		perSection: perSection,
	}

	for i := 0; i < section*section; i++ {
		m = m.randomizeSection(i)
	}

	return m, nil
}
