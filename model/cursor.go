package model

func (m model) cursorCellUp() {
	if m.cursor.y >= 1 {
		m.cursor.y--
	}
}

func (m model) cursorCellDown() {
	if m.cursor.y < m.fieldSize-1 {
		m.cursor.y++
	}
}

func (m model) cursorCellRight() {
	if m.cursor.x < m.fieldSize-1 {
		m.cursor.x++
	}
}

func (m model) cursorCellLeft() {
	if m.cursor.x >= 1 {
		m.cursor.x--
	}
}

func (m model) cursorSectionUp() {
	if m.cursor.y >= m.cellsPerSection {
		m.cursor.y -= m.cellsPerSection
	}
}

func (m model) cursorSectionDown() {
	if m.cursor.y < m.fieldSize-m.cellsPerSection {
		m.cursor.y += m.cellsPerSection
	}
}

func (m model) cursorSectionRight() {
	if m.cursor.x < m.fieldSize-m.cellsPerSection {
		m.cursor.x += m.cellsPerSection
	}
}

func (m model) cursorSectionLeft() {
	if m.cursor.x >= m.cellsPerSection {
		m.cursor.x -= m.cellsPerSection
	}
}
