package state

import (
	"fmt"
	"strconv"
)

// History returns the gameplay history for replays
func (s *State) History() string {
	return s.history
}

func UnmarshalAll(history string) (*State, error) {
	s := Unmarshal(history)
	for {
		res, err := s.Step()
		if err != nil {
			return nil, err
		}
		if !res {
			break
		}
	}
	return s, nil
}

func Unmarshal(history string) *State {
	return &State{
		history:      history,
		historyIndex: 0,
		cursor:       &Point{0, 0},
	}
}

func (s *State) Step() (bool, error) {
	if s.historyIndex >= len(s.history) {
		return false, nil
	}

	switch s.history[s.historyIndex] {
	case 'i', 'm', 'f', 'c', 'r', 'l', 'u', 'd', 'R', 'L', 'U', 'D':
		err := s.beforeAct()
		if err != nil {
			return false, err
		}
	}

	switch s.history[s.historyIndex] {
	case 'g':
		if s.initialized {
			return false, fmt.Errorf("initialization step in invalid location (index %d)", s.historyIndex)
		}
		err := s.historyStart()
		if err != nil {
			return false, err
		}
	case 'i':
		err := s.historyInitSection()
		if err != nil {
			return false, err
		}
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
	case 't': // TODO for now, this is just sugar. Will be a timestamp for events completed.
		for {
			s.historyIndex++
			if s.historyIndex >= len(s.history) || s.history[s.historyIndex] == ')' {
				break
			}
		}
	case ' ', '\n', '\t':
		break
	case '#':
		for {
			s.historyIndex++
			if s.historyIndex >= len(s.history) || s.history[s.historyIndex] == '\n' {
				break
			}
		}
	default:
		return false, fmt.Errorf("invalid step in history: %s (index %d)", string(s.history[s.historyIndex]), s.historyIndex)
	}
	s.historyIndex++

	return true, nil
}

func (s *State) beforeAct() error {
	if !s.initialized {
		return fmt.Errorf("tried to act on an uninitialized state (index %d)", s.historyIndex)
	}
	return nil
}

func (s *State) historyStart() error {
	peek := s.historyIndex + 1
	p, peek, err := s.historyPoint(peek)
	if err != nil {
		return err
	}
	s.sectionSize = p.X
	s.cellsPerSection = p.Y
	s.cells = newField(s.size())
	s.sections = newField(s.sectionSize)
	s.cursor = &Point{0, 0}
	s.rowHeaders = make([]header, s.size())
	s.colHeaders = make([]header, s.size())
	s.historyIndex = peek
	s.score.Blackout = make([]bool, s.size())
	s.initialized = true
	return nil
}

func (s *State) historyInitSection() error {
	peek := s.historyIndex + 1
	p, peek, err := s.historyPoint(peek)
	if err != nil {
		return err
	}

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
	return nil
}

func (s *State) historyPoint(index int) (Point, int, error) {
	var x, y string

	// Advance past opening paren
	index++
	for {
		if index >= len(s.history) {
			return Point{}, 0, fmt.Errorf("point.X never ended? (index %d)", index)
		}
		switch s.history[index] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			x += string(s.history[index])
			index++
			continue
		case ',':
			index++
		default:
			return Point{}, 0, fmt.Errorf("invalid character in point.X: %s (index %d)", string(s.history[index]), index)
		}
		break
	}

	for {
		if index >= len(s.history) {
			return Point{}, 0, fmt.Errorf("point.Y never ended? (index %d)", index)
		}
		switch s.history[index] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			y += string(s.history[index])
			index++
			continue
		case ')':
			break
		default:
			return Point{}, 0, fmt.Errorf("invalid character in point.Y: %s (index %d)", string(s.history[index]), index)
		}
		break
	}

	pX, _ := strconv.Atoi(x)
	pY, _ := strconv.Atoi(y)
	return Point{pX, pY}, index, nil
}
