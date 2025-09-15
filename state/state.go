package state

import (
	"bytes"
	"fmt"
	"math/rand"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

type State struct {
	cursor *Point

	clears, score, factor int
	blackout              []bool

	history      string
	historyIndex int

	sectionSize, cellsPerSection int
	cells, sections              *field
}

func New(sectionSize, cellsPerSection int) *State {
	s := &State{
		sectionSize:     sectionSize,
		cellsPerSection: cellsPerSection,
		cells:           newField(sectionSize * cellsPerSection),
		sections:        newField(sectionSize),
		cursor:          &Point{0, 0},
	}
	s.history = fmt.Sprintf("g(%d,%d)", sectionSize, cellsPerSection)
	for x := 0; x < sectionSize; x++ {
		for y := 0; y < sectionSize; y++ {
			s.initSection(Point{x, y})
		}
	}
	// TODO reveal 1 row per section, 1 col per sections/2
	return s
}

// size is a utility method to keep cursor code clean.
func (s *State) size() int {
	return s.sectionSize * s.cellsPerSection
}

func (s *State) String() string {
	var buf bytes.Buffer
	for x := 0; x < s.sectionSize*s.cellsPerSection; x++ {
		for y := 0; y < s.sectionSize*s.cellsPerSection; y++ {
			p := Point{x, y}
			if s.cells.correct(p) && (s.cells.marked(p) || s.cells.flagged(p)) {
				if s.cells.state(p) {
					buf.WriteString("O")
				} else {
					buf.WriteString("X")
				}
			} else {
				if s.cells.marked(p) {
					buf.WriteString("o")
				} else if s.cells.flagged(p) {
					buf.WriteString("x")
				} else {
					if s.cells.state(p) {
						buf.WriteString(".")
					} else {
						buf.WriteString(" ")
					}
				}
			}
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func (s *State) Mark() []bool {
	s.history += "m"
	return s.mark(*s.cursor)
}

func (s *State) Flag() []bool {
	s.history += "f"
	return s.flag(*s.cursor)
}

func (s *State) Clear() {
	s.history += "c"
	s.clear(*s.cursor, true)
}

func (s *State) view(cursor []int) {
}

func (s *State) initSection(p Point) {
	s.history += fmt.Sprintf("i%s", p)
	s.sections.clear(p, false)
	startX := p.X * s.cellsPerSection
	startY := p.Y * s.cellsPerSection
	for x := 0; x < s.cellsPerSection; x++ {
		for y := 0; y < s.cellsPerSection; y++ {
			currPoint := Point{x + startX, y + startY}
			s.cells.clear(currPoint, true)
			if rand.Int()%2 == 0 {
				s.cells.kill(currPoint)
				s.history += "x"
			} else {
				s.cells.vivify(currPoint)
				s.history += "o"
			}
		}
	}
	s.update(p)
}

func (s *State) mark(p Point) []bool {
	s.cells.mark(p)
	return s.update(p)
}

func (s *State) flag(p Point) []bool {
	s.cells.flag(p)
	return s.update(p)
}

func (s *State) clear(p Point, deadCorrect bool) {
	s.cells.clear(p, deadCorrect)
	s.update(p)
}

func (s *State) sectionCorrect(p Point) bool {
	sectionCorrect := true
	for x := 0; x < s.cellsPerSection; x++ {
		for y := 0; y < s.cellsPerSection; y++ {
			sectionCorrect = sectionCorrect && s.cells.correct(Point{p.X*s.cellsPerSection + x, p.Y*s.cellsPerSection + y})
		}
	}
	return sectionCorrect
}

func (s *State) updateCompletedSections() {
	for x := 0; x < s.sectionSize; x++ {
		for y := 0; y < s.sectionSize; y++ {
			complete := true
			for i := 0; i < s.sectionSize; i++ {
				complete = complete &&
					s.sections.correct(Point{x, i}) &&
					s.sections.correct(Point{i, y})
			}
			s.sections.setComplete(Point{x, y}, complete)
		}
	}
}

func (s *State) clearValidCompletedSections() []bool {
	cleared := make([]bool, s.sections.size*s.sections.size)
	for x := 0; x < s.sectionSize; x++ {
		for y := 0; y < s.sectionSize; y++ {
			clearable := s.sections.complete(Point{x, y}) &&
				s.sections.complete(Point{(x + 1) % s.sectionSize, y}) &&
				s.sections.complete(Point{x, (y + 1) % s.sectionSize}) &&
				s.sections.complete(Point{(x + 1) % s.sectionSize, (y + 1) % s.sectionSize})
			if clearable {
				cleared[y*s.sectionSize+x] = true
				cleared[(y*s.sectionSize + (x+1)%s.sectionSize)] = true
				cleared[((y*s.sectionSize+s.sectionSize)%(len(cleared)) + x)] = true
				cleared[((y*s.sectionSize+s.sectionSize)+(x+1%s.sectionSize))%len(cleared)] = true
			}
		}
	}
	for i, v := range cleared {
		if v {
			s.initSection(Point{i % s.sectionSize, i / s.sectionSize})
		}
	}
	return cleared
}

func (s *State) update(p Point) []bool {
	sectionPoint := Point{p.X / s.cellsPerSection, p.Y / s.cellsPerSection}

	if s.cells.correct(p) && s.sectionCorrect(sectionPoint) {
		s.sections.setCorrect(sectionPoint, true)
		s.updateCompletedSections()
		return s.clearValidCompletedSections()
	}
	s.sections.setCorrect(sectionPoint, false)
	for i := 0; i < s.sectionSize; i++ {
		s.sections.setComplete(Point{sectionPoint.X, i}, false)
		s.sections.setComplete(Point{i, sectionPoint.Y}, false)
	}
	return make([]bool, s.sections.size*s.sections.size)
}
