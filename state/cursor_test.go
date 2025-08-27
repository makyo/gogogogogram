package state

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCursor(t *testing.T) {
	Convey("Given a cursor", t, func() {
		s := New(4, 4)
		So(*s.cursor, ShouldResemble, Point{0, 0})

		Convey("When moving cell to cell", func() {

			Convey("You can move down", func() {
				s.CursorCellDown()
				So(*s.cursor, ShouldResemble, Point{0, 1})
			})

			Convey("You can move up", func() {
				s.CursorCellDown()
				s.CursorCellUp()
				So(*s.cursor, ShouldResemble, Point{0, 0})
			})

			Convey("You can move right", func() {
				s.CursorCellRight()
				So(*s.cursor, ShouldResemble, Point{1, 0})
			})

			Convey("You can move left", func() {
				s.CursorCellRight()
				s.CursorCellLeft()
				So(*s.cursor, ShouldResemble, Point{0, 0})
			})

			Convey("You can't move up beyond the top", func() {
				s.CursorCellUp()
				So(*s.cursor, ShouldResemble, Point{0, 0})
			})

			Convey("You can't move left beyond the edge", func() {
				s.CursorCellLeft()
				So(*s.cursor, ShouldResemble, Point{0, 0})
			})

			Convey("You can't move down below the bottom", func() {
				s.cursor = &Point{15, 15}
				s.CursorCellDown()
				So(*s.cursor, ShouldResemble, Point{15, 15})
			})

			Convey("You can't move right beyond the edge", func() {
				s.cursor = &Point{15, 15}
				s.CursorCellRight()
				So(*s.cursor, ShouldResemble, Point{15, 15})
			})
		})

		Convey("When moving section to section", func() {
			Convey("You can move down", func() {
				s.CursorSectionDown()
				So(*s.cursor, ShouldResemble, Point{0, 4})
			})

			Convey("You can move up", func() {
				s.CursorSectionDown()
				s.CursorSectionUp()
				So(*s.cursor, ShouldResemble, Point{0, 0})
			})

			Convey("You can move right", func() {
				s.CursorSectionRight()
				So(*s.cursor, ShouldResemble, Point{4, 0})
			})

			Convey("You can move left", func() {
				s.CursorSectionRight()
				s.CursorSectionLeft()
				So(*s.cursor, ShouldResemble, Point{0, 0})
			})

			Convey("You can't move up beyond the top", func() {
				s.CursorCellUp()
				So(*s.cursor, ShouldResemble, Point{0, 0})
			})

			Convey("You can't move left beyond the edge", func() {
				s.CursorCellLeft()
				So(*s.cursor, ShouldResemble, Point{0, 0})
			})

			Convey("You can't move down below the bottom", func() {
				s.cursor = &Point{15, 15}
				s.CursorSectionDown()
				So(*s.cursor, ShouldResemble, Point{15, 15})
			})

			Convey("You can't move right beyond the edge", func() {
				s.cursor = &Point{15, 15}
				s.CursorSectionRight()
				So(*s.cursor, ShouldResemble, Point{15, 15})
			})
		})
	})
}
