package state

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestState(t *testing.T) {
	Convey("Given a game state", t, func() {
		s := New(2, 2)

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
				So(s.String(), ShouldEqual, "OX  \n    \n    \n    \n")
			})

			Convey("It shows incorrect guesses", func() {
				s.Flag()
				s.mark(Point{0, 1})
				So(s.String(), ShouldEqual, "xo  \n    \n    \n    \n")
			})

			Convey("It shows cleared guesses as empty", func() {
				s.Flag()
				s.mark(Point{0, 1})
				s.Clear()
				So(s.String(), ShouldEqual, " o  \n    \n    \n    \n")
			})
		})
	})
}
