package state

func (s *State) CursorCellUp() {
	if s.cursorCellUp() {
		s.history += "u"
	}
}

func (s *State) CursorCellDown() {
	if s.cursorCellDown() {
		s.history += "d"
	}
}

func (s *State) CursorCellRight() {
	if s.cursorCellRight() {
		s.history += "r"
	}
}

func (s *State) CursorCellLeft() {
	if s.cursorCellLeft() {
		s.history += "l"
	}
}

func (s *State) CursorSectionUp() {
	if s.cursorSectionUp() {
		s.history += "U"
	}
}

func (s *State) CursorSectionDown() {
	if s.cursorSectionDown() {
		s.history += "D"
	}
}

func (s *State) CursorSectionRight() {
	if s.cursorSectionRight() {
		s.history += "R"
	}
}

func (s *State) CursorSectionLeft() {
	if s.cursorSectionLeft() {
		s.history += "L"
	}
}

func (s *State) cursorCellUp() bool {
	if s.cursor.Y >= 1 {
		s.cursor.Y--
		return true
	}
	return false
}

func (s *State) cursorCellDown() bool {
	if s.cursor.Y < s.size()-1 {
		s.cursor.Y++
		return true
	}
	return false
}

func (s *State) cursorCellRight() bool {
	if s.cursor.X < s.size()-1 {
		s.cursor.X++
		return true
	}
	return false
}

func (s *State) cursorCellLeft() bool {
	if s.cursor.X >= 1 {
		s.cursor.X--
		return true
	}
	return false
}

func (s *State) cursorSectionUp() bool {
	if s.cursor.Y >= s.cellsPerSection {
		s.cursor.Y -= s.cellsPerSection
		return true
	}
	return false
}

func (s *State) cursorSectionDown() bool {
	if s.cursor.Y < s.size()-s.cellsPerSection {
		s.cursor.Y += s.cellsPerSection
		return true
	}
	return false
}

func (s *State) cursorSectionRight() bool {
	if s.cursor.X < s.size()-s.cellsPerSection {
		s.cursor.X += s.cellsPerSection
		return true
	}
	return false
}

func (s *State) cursorSectionLeft() bool {
	if s.cursor.X >= s.cellsPerSection {
		s.cursor.X -= s.cellsPerSection
		return true
	}
	return false
}
