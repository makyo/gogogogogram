package state

import (
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestState(t *testing.T) {
	Convey("Given a game state", t, func() {
		s := New(2, 2)
		for x := 0; x < 4; x++ {
			for y := 0; y < 4; y++ {
				s.cells.kill(Point{x, y})
				if x < 2 && y < 2 {
					s.sections.clear(Point{x, y}, false)
				}
			}
		}

		Convey("You can get the size", func() {
			So(s.size(), ShouldEqual, 4)
		})

		Convey("You can render a simple string of the board", func() {
			s.cells.vivify(Point{0, 0})
			s.cells.kill(Point{0, 1})
			s.cells.vivify(Point{1, 0})
			s.cells.kill(Point{1, 1})

			Convey("It shows correct guesses", func() {
				s.Mark()
				s.flag(Point{0, 1})
				So(s.String(), ShouldEqual, "OX  \n.   \n    \n    \n")
			})

			Convey("It shows incorrect guesses", func() {
				s.Flag()
				s.mark(Point{0, 1})
				So(s.String(), ShouldEqual, "xo  \n.   \n    \n    \n")
			})

			Convey("It shows cleared guesses as empty", func() {
				s.Flag()
				s.mark(Point{0, 1})
				s.Clear()
				So(s.String(), ShouldEqual, ".o  \n.   \n    \n    \n")
			})
		})

		Convey("You can maintain a history of all actions", func() {
			s.Flag()
			s.Mark()
			s.Clear()
			s.CursorSectionRight()
			s.CursorSectionLeft()
			s.CursorSectionDown()
			s.CursorSectionUp()
			s.CursorCellRight()
			s.CursorCellLeft()
			s.CursorCellDown()
			s.CursorCellUp()

			matches, err := regexp.Match("gogogogogram 2 2:\n(i\\(\\d,\\d\\)[\b\x01]{4}){4}fmcRLDUrldu", []byte(s.History()))
			So(matches, ShouldBeTrue)
			So(err, ShouldBeNil)
		})

		Convey("You can complete sections and clear portions of the board", func() {
			s.cells.vivify(Point{0, 0})
			s.cells.vivify(Point{2, 0})
			s.cells.vivify(Point{0, 2})
			s.cells.vivify(Point{2, 2})

			Convey("It marks sections as correct when they are correctly guessed", func() {
				res := s.mark(Point{0, 0})
				So(s.cells.correct(Point{0, 0}), ShouldBeTrue)
				So(s.sections.correct(Point{0, 0}), ShouldBeTrue)
				So(s.String(), ShouldEqual, "O . \n    \n. . \n    \n")
				So(res, ShouldResemble, make([]bool, 4))
			})

			Convey("It marks sections as complete if they meet the criteria", func() {
				s.mark(Point{0, 0})
				s.mark(Point{2, 0})
				res := s.mark(Point{0, 2})
				So(res, ShouldResemble, make([]bool, 4))
				So(s.String(), ShouldEqual, "O O \n    \nO . \n    \n")
				So(s.sections.correct(Point{0, 0}), ShouldBeTrue)
				So(s.sections.correct(Point{1, 0}), ShouldBeTrue)
				So(s.sections.correct(Point{0, 1}), ShouldBeTrue)
				So(s.sections.complete(Point{0, 0}), ShouldBeTrue)
			})

			Convey("It clears sections when they meet the criteria", func() {
				s.mark(Point{0, 0})
				s.mark(Point{2, 0})
				s.mark(Point{0, 2})
				res := s.mark(Point{2, 2})
				So(res, ShouldResemble, []bool{true, true, true, true})
			})
		})
	})
}
