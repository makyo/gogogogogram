package model

import "fmt"

func (m Model) Mark() Model {
	if m.marks&m.cursor != 0 {
		m.marks |= m.cursor
		m.flags = m.flags &^ m.cursor
	} else {
		m.marks = m.marks &^ m.cursor
	}
	m.history = append(m.history, fmt.Sprintf("m%d", m.cursor))
	m = m.update()
	return m
}

func (m Model) Flag() Model {
	if m.flags&m.cursor != 0 {
		m.flags |= m.cursor
		m.marks = m.marks &^ m.cursor
	} else {
		m.flags = m.flags &^ m.cursor
	}
	m.history = append(m.history, fmt.Sprintf("f%d", m.cursor))
	m = m.update()
	return m
}

func (m Model) ClearGuess() Model {
	m.marks = m.marks &^ m.cursor
	m.flags = m.flags &^ m.cursor
	m.correct = m.correct &^ m.cursor
	m.history = append(m.history, fmt.Sprintf("c%d", m.cursor))
	m = m.update()
	return m
}
