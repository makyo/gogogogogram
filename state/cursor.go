package state

func (s *State) CursorCellUp() {
	if s.cursor.Y >= 1 {
		s.cursor.Y--
		s.history += "u"
	}
}

func (s *State) CursorCellDown() {
	if s.cursor.Y < s.size()-1 {
		s.cursor.Y++
		s.history += "d"
	}
}

func (s *State) CursorCellRight() {
	if s.cursor.X < s.size()-1 {
		s.cursor.X++
		s.history += "r"
	}
}

func (s *State) CursorCellLeft() {
	if s.cursor.X >= 1 {
		s.cursor.X--
		s.history += "l"
	}
}

func (s *State) CursorSectionUp() {
	if s.cursor.Y >= s.cellsPerSection {
		s.cursor.Y -= s.cellsPerSection
		s.history += "U"
	}
}

func (s *State) CursorSectionDown() {
	if s.cursor.Y < s.size()-s.cellsPerSection {
		s.cursor.Y += s.cellsPerSection
		s.history += "D"
	}
}

func (s *State) CursorSectionRight() {
	if s.cursor.X < s.size()-s.cellsPerSection {
		s.cursor.X += s.cellsPerSection
		s.history += "R"
	}
}

func (s *State) CursorSectionLeft() {
	if s.cursor.X >= s.cellsPerSection {
		s.cursor.X -= s.cellsPerSection
		s.history += "L"
	}
}
