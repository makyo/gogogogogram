package state

import "strconv"

// History returns the gameplay history for replays
func (s *State) History() string {
	return s.history
}

func UnmarshalAll(history string) *State {
	s := Unmarshal(history)
	for {
		if !s.Step() {
			break
		}
	}
	return s
}

func Unmarshal(history string) *State {
	return &State{
		history:      history,
		historyIndex: 0,
		cursor:       &Point{0, 0},
	}
}

func (s *State) Step() bool {
	if s.historyIndex >= len(s.history) {
		return false
	}

	switch s.history[s.historyIndex] {
	case 'g':
		s.historyStart()
	case 'i':
		s.historyInitSection()
	case 'm':
		s.mark(*s.cursor)
	case 'f':
		s.flag(*s.cursor)
	case 'c':
		s.clear(*s.cursor, true)
	case 'r':
		s.cursorCellRight()
	case 'l':
		s.cursorCellLeft()
	case 'u':
		s.cursorCellUp()
	case 'd':
		s.cursorCellDown()
	case 'R':
		s.cursorSectionRight()
	case 'L':
		s.cursorSectionLeft()
	case 'U':
		s.cursorSectionUp()
	case 'D':
		s.cursorSectionDown()
	}
	s.historyIndex++

	return true
}

func (s *State) historyStart() {
	peek := s.historyIndex + 1
	p, peek := s.historyPoint(peek)
	s.sectionSize = p.X
	s.cellsPerSection = p.Y
	s.cells = newField(s.size())
	s.sections = newField(s.sectionSize)
	s.cursor = &Point{0, 0}
	s.rowHeaders = make([]header, s.size())
	s.colHeaders = make([]header, s.size())
	s.historyIndex = peek
	s.score.Blackout = make([]bool, s.size())
}

func (s *State) historyInitSection() {
	peek := s.historyIndex + 1
	p, peek := s.historyPoint(peek)
	segment := []byte(s.history[peek : peek+s.cellsPerSection*s.cellsPerSection])
	for i, c := range segment {
		curr := Point{
			p.X*s.cellsPerSection + (i % s.cellsPerSection),
			p.Y*s.cellsPerSection + (i / s.cellsPerSection),
		}
		s.cells.clear(curr, true)
		switch c {
		case 'x':
			s.cells.vivify(curr)
		case 'o':
			s.cells.kill(curr)
		default:
			break
		}
		peek++
	}
	s.historyIndex = peek
}

func (s *State) historyPoint(index int) (Point, int) {
	var x, y string

	// Advance past paren
	index++
	for {
		switch s.history[index] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			x += string(s.history[index])
			index++
			continue
		}
		break
	}

	// Advance past comma
	index++
	for {
		switch s.history[index] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			y += string(s.history[index])
			index++
			continue
		}
		break
	}

	pX, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	pY, err := strconv.Atoi(y)
	if err != nil {
		panic(err)
	}
	return Point{pX, pY}, index
}
