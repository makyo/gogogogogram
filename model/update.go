package model

func (m Model) update() Model {
	// Reset sections to all true
	m.sections = m.section*m.section - 1

	// Update correctness/sections
	for i := 0; i < m.size*m.size; i++ {
		cell := m.field & i
		marked := m.marks & i
		flagged := m.flags & i

		// Update the correctness of each cell
		m.correct |= (cell & marked) | (flagged &^ cell)

		// Update the correctness of each section
		cellSection := ((i / (m.section * m.size)) * m.section) + (i/m.perSection)%m.section
		m.sections &= m.correct & cellSection
	}

	// Check for complete sections, which are those where the row and column are both correct
	completedIndices := map[int]bool{}
	for i := 0; i < m.section*m.section; i++ {
		if m.sections&i == 0 {
			continue
		}
		mask := 0
		x := i % m.section
		y := i / m.section
		for j := 0; j < m.section; j++ {
			mask |= 1<<x + (j * m.section)
			mask |= 1<<y + j
		}
		if m.sections&mask == mask {
			completedIndices[i] = true
		}
	}

	// Check for clears
	clearedIndices := map[int]bool{}
	clearedIncrease := 0
	for i, _ := range completedIndices {
		x := i % m.section
		y := i / m.section

		right := x + 1
		if right == m.section {
			right = -1
		}

		down := m.section
		if y+1 == m.section {
			down = -m.section
		}

		_, okRight := completedIndices[i+right]
		_, okDown := completedIndices[i+down]
		_, okDownRight := completedIndices[i+right+down]
		if okRight && okDown && okDownRight {
			clearedIndices[i] = true
			clearedIndices[i+right] = true
			clearedIndices[i+down] = true
			clearedIndices[i+down+right] = true
			clearedIncrease++
			m.completed |= 1 << i
			m.completed |= 1<<i + right
			m.completed |= 1<<i + down
			m.completed |= 1<<i + down + right
		}
	}

	// Clear and bump scores
	for i, _ := range clearedIndices {
		m = m.randomizeSection(i)
	}
	if clearedIncrease > 0 {
		m.score += clearedIncrease * (m.factor + 1)
		m.clears++
	}

	// Check for blackout
	if m.completed == m.section*m.section-1 {
		m.completed = 0
		m.factor++
	}

	// Update row/column states/correctness
	for x := 0; x < m.size; x++ {
		rowCorrect := true
		columnCorrect := true
		for y := 0; y < m.size; y++ {
			cellCorrect := m.correct&y*m.size+x != 0
			rowCorrect = rowCorrect && cellCorrect
			columnCorrect = columnCorrect && cellCorrect
		}
	}

	return m
}
