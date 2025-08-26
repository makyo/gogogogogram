package model

func (m Model) CursorCellUp() Model {
	if m.cursor-m.size > 0 {
		m.cursor -= m.size
	}
	return m
}

func (m Model) CursorCellDown() Model {
	if m.cursor+m.size < m.size*m.size {
		m.cursor += m.size
	}
	return m
}

func (m Model) CursorCellRight() Model {
	if m.cursor%m.size < m.size {
		m.cursor++
	}
	return m
}

func (m Model) CursorCellLeft() Model {
	if m.cursor%m.size != 1 {
		m.cursor--
	}
	return m
}

func (m Model) CursorSectionUp() Model {
	if m.cursor > m.size*m.perSection {
		m.cursor -= m.size * m.perSection
	}
	return m
}

func (m Model) CursorSectionDown() Model {
	if m.cursor < m.size*m.perSection*(m.section-1) {
		m.cursor += m.size * m.perSection
	}
	return m
}

func (m Model) CursorSectionRight() Model {
	if m.cursor%m.perSection < m.section-1 {
		m.cursor += m.section
	}
	return m
}

func (m Model) CursorSectionLeft() Model {
	if m.cursor%m.perSection > 0 {
		m.cursor -= m.section
	}
	return m
}
