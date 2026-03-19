package state

import (
	"errors"
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHistory(t *testing.T) {
	Convey("Given a game's history", t, func() {

		Convey("You can maintain a history of all actions", func() {
			s := New(2, 2)
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

			matches, err := regexp.Match(`g\(2,2\)(i\(\d,\d\)[xo]{4}){4}fmcRLDUrldu`, []byte(s.History()))

			So(err, ShouldBeNil)
			So(matches, ShouldBeTrue)
		})

		Convey("You can load a game from its saved history", func() {
			history := "g(2,2)i(0,0)oxoxi(1,0)xoxoi(0,1)ooooi(1,1)xxxx\n# Here we goooo~\nfRmDcLUrldut(1773881959)"
			s, err := UnmarshalAll(history)
			So(err, ShouldBeNil)

			Convey("It sets the history", func() {
				So(s.History(), ShouldEqual, history)
			})

			Convey("It correctly sets the field sizes", func() {
				So(s.sectionSize, ShouldEqual, 2)
				So(s.cellsPerSection, ShouldEqual, 2)
				So(len(s.cells.cells), ShouldEqual, 16)
				So(len(s.sections.cells), ShouldEqual, 4)
			})

			Convey("The cursor should be set to the appropriate coordinates after all movement", func() {
				So(s.cursor, ShouldResemble, &Point{0, 0})
			})

			Convey("The board should be appropriately marked and filled", func() {
				So(s.String(), ShouldEqual, "X.  \n    \no  .\n....\n")
			})
		})

		Convey("Errors are handled during loading", func() {
			Convey("It raises errors if acting on an uninitialized state", func() {
				s, err := UnmarshalAll("m")
				So(err, ShouldResemble, errors.New("tried to act on an uninitialized state (index 0)"))
				So(s, ShouldBeNil)
			})

			Convey("It raises errors if trying to initialize an initialized state", func() {
				s, err := UnmarshalAll("g(1,1)g(1,1)")
				So(err, ShouldResemble, errors.New("initialization step in invalid location (index 6)"))
				So(s, ShouldBeNil)
			})

			Convey("It raises errors in reading points", func() {
				s, err := UnmarshalAll("g(1")
				So(err, ShouldResemble, errors.New("point.X never ended? (index 3)"))
				So(s, ShouldBeNil)

				s, err = UnmarshalAll("g(a,1)")
				So(err, ShouldResemble, errors.New("invalid character in point.X: a (index 2)"))
				So(s, ShouldBeNil)

				s, err = UnmarshalAll("g(1,1)i(0,1")
				So(err, ShouldResemble, errors.New("point.Y never ended? (index 11)"))
				So(s, ShouldBeNil)

				s, err = UnmarshalAll("g(1,a)")
				So(err, ShouldResemble, errors.New("invalid character in point.Y: a (index 4)"))
				So(s, ShouldBeNil)
			})

			Convey("It only accepts valid steps", func() {
				history := "g(2,2)i(0,0)oxoxi(1,0)xoxoi(0,1)ooooi(1,1)xxxxz"
				s, err := UnmarshalAll(history)
				So(err, ShouldResemble, errors.New("invalid step in history: z (index 46)"))
				So(s, ShouldBeNil)
			})
		})
	})
}
