package state

func (s *State) CursorCellUp() {
	if s.cursor.Y >= 1 {
		s.cursor.Y--
	}
}

func (s *State) CursorCellDown() {
	if s.cursor.Y < s.size()-1 {
		s.cursor.Y++
	}
}

func (s *State) CursorCellRight() {
	if s.cursor.X < s.size()-1 {
		s.cursor.X++
	}
}

func (s *State) CursorCellLeft() {
	if s.cursor.X >= 1 {
		s.cursor.X--
	}
}

func (s *State) CursorSectionUp() {
	if s.cursor.Y >= s.cellsPerSection {
		s.cursor.Y -= s.cellsPerSection
	}
}

func (s *State) CursorSectionDown() {
	if s.cursor.Y < s.size()-s.cellsPerSection {
		s.cursor.Y += s.cellsPerSection
	}
}

func (s *State) CursorSectionRight() {
	if s.cursor.X < s.size()-s.cellsPerSection {
		s.cursor.X += s.cellsPerSection
	}
}

func (s *State) CursorSectionLeft() {
	if s.cursor.X >= s.cellsPerSection {
		s.cursor.X -= s.cellsPerSection
	}
}
