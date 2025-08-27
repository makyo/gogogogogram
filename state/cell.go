package state

const (
	statebit    = 1             // the current state of the cell (for field)
	flagbit     = statebit << 1 // whether or not the cell has been flagged as empty (for field)
	markbit     = statebit << 2 // whether or not the cell has been flagged as full (for field)
	correctbit  = statebit << 3 // whether or not the cell is correct (for field, section)
	completebit = statebit << 4 // whether or not the cell is complete (for section)
)

// Cells {{{

// cell represents a single entry in the field managing various states.
type cell byte

// state returns whether the cell is alive or dead.
func (c cell) state() bool {
	return (c & statebit) != 0
}

// correct returns whether or not the cell has been guessed correctly, or the section is made up of entirely correct guesses.
func (c cell) correct() bool {
	return (c & correctbit) != 0
}

// complete marks the section as complete (that is, all sections in its row/column are correct.
func (c cell) complete() bool {
	return (c & completebit) != 0
}

// flagged returns whether or not the cell is suspected to be dead.
func (c cell) flagged() bool {
	return (c & flagbit) != 0
}

// marked returns whether or not the cell is suspected to be alive.
func (c cell) marked() bool {
	return (c & markbit) != 0
}

// vivify sets the state of the cell to alive.
func (c cell) vivify() cell {
	c = c | statebit
	if c.marked() {
		return c.setCorrect(true)
	} else {
		return c.setCorrect(false)
	}
}

// kill sets the state of the cell to dead.
func (c cell) kill() cell {
	c = c &^ statebit
	if c.flagged() {
		return c.setCorrect(true)
	} else {
		return c.setCorrect(false)
	}
}

// setCorrect marks the cell as being guessed correctly, or the section being made up of entirely correct guesses.
func (c cell) setCorrect(to bool) cell {
	if to {
		return c | correctbit
	} else {
		return c &^ correctbit
	}
}

// setComplete marks the section as complete (that is, all sections in its row/column are correct).
func (c cell) setComplete(to bool) cell {
	if to {
		return c | completebit
	} else {
		return c &^ completebit
	}
}

// mark marks a bit as suspected alive.
func (c cell) mark() cell {
	c |= markbit
	c = c &^ flagbit
	if c.state() {
		return c.setCorrect(true)
	} else {
		return c.setCorrect(false)
	}
}

// flag marks a bit as suspected dead.
func (c cell) flag() cell {
	c |= flagbit
	c = c &^ markbit
	if c.state() {
		return c.setCorrect(false)
	} else {
		return c.setCorrect(true)
	}
}

// clear clears all bits except for the status.
func (c cell) clear() cell {
	return c &^ (markbit | flagbit | correctbit | completebit)
}
