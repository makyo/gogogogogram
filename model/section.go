package model

import "math/rand"

func (m Model) randomizeSection(s int) Model {
	m = m.clearSection(s)
	for y := 0; y < m.section; y++ {
		for x := 0; x < m.section; x++ {
			cell := ((s/m.section)*m.size + y) + ((s%m.size)*m.section + x)
			if rand.Int()%2 == 1 {
				m.field |= cell
			}
		}
	}
	return m
}

func (m Model) clearSection(s int) Model {
	m.sections = m.sections &^ s
	for y := 0; y < m.section; y++ {
		for x := 0; x < m.section; x++ {
			cell := ((s/m.section)*m.size + y) + ((s%m.size)*m.section + x)
			m.field = m.field &^ cell
			m.marks = m.marks &^ cell
			m.flags = m.flags &^ cell
			m.correct = m.correct &^ cell
		}
	}
	return m
}
