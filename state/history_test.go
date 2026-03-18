package state

import (
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
			history := "g(2,2)i(0,0)oxoxi(1,0)xoxoi(0,1)ooooi(1,1)xxxxfRmDcLUrldu"
			s := UnmarshalAll(history)

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
	})
}
