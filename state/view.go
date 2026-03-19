package state

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

var (
	complete = []string{"██ ", "██ "}
	flagged  = []string{"╲╱ ", "╱╲ "}
	marked   = []string{"╭╮ ", "╰╯ "}
	blank    = []string{".- ", "   "}

	cursorStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#005566"))

	gridStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Black)
)

func (c *cell) View() []string {
	if c.complete() {
		return complete
	} else if c.flagged() {
		return flagged
	} else if c.marked() {
		return marked
	} else {
		return blank
	}
}

func (f *field) View(p *Point, sectionSize, cellsPerSection int) string {
	res := gridStyle.Render("┌")
	for c := 0; c < cellsPerSection*sectionSize*3; c++ {
		res += gridStyle.Render("─")
		if c == cellsPerSection*sectionSize*3-1 {
			res += gridStyle.Render("┐") + "\n"
			break
		}
		if c%(cellsPerSection*3) == cellsPerSection*3-1 {
			res += gridStyle.Render("┬")
		}
	}
	resA := gridStyle.Render("│")
	resB := gridStyle.Render("│")
	for i, c := range f.cells {
		cellRes := c.View()
		if i == f.i(*p) {
			resA += cursorStyle.Render(cellRes[0])
			resB += cursorStyle.Render(cellRes[1])
		} else {
			resA += cellRes[0]
			resB += cellRes[1]
		}
		if i%cellsPerSection == cellsPerSection-1 {
			resA += gridStyle.Render("│")
			resB += gridStyle.Render("│")
		}
		if i%f.size == f.size-1 {
			res += fmt.Sprintf("%s\n%s\n", resA, resB)
			resA = gridStyle.Render("│")
			resB = gridStyle.Render("│")
		}
		if i != 0 && i%(cellsPerSection*cellsPerSection*sectionSize) == 0 {
			res += gridStyle.Render("├")
			for section := 0; section < sectionSize; section++ {
				for sectionCell := 0; sectionCell < cellsPerSection; sectionCell++ {
					res += gridStyle.Render("───")
				}
				if section < sectionSize-1 {
					res += gridStyle.Render("┼")
				}
			}
			res += gridStyle.Render("┤") + "\n"
		}
	}
	res += gridStyle.Render("└")
	for c := 0; c < cellsPerSection*sectionSize*3; c++ {
		res += gridStyle.Render("─")
		if c == cellsPerSection*sectionSize*3-1 {
			res += gridStyle.Render("┘\n")
			break
		}
		if c%(cellsPerSection*3) == cellsPerSection*3-1 {
			res += gridStyle.Render("┴")
		}
	}
	return res
}

func (s *State) View() string {
	return s.cells.View(s.cursor, s.sectionSize, s.cellsPerSection)
}
