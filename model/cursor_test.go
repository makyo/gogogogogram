package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCursor(t *testing.T) {
	Convey("Given a cursor", t, func() {
		m := New(4, 4)
		So(*m.cursor, ShouldResemble, point{0, 0})

		Convey("When moving cell to cell", func() {

			Convey("You can move down", func() {
				m.cursorCellDown()
				So(*m.cursor, ShouldResemble, point{0, 1})
			})

			Convey("You can move up", func() {
				m.cursorCellDown()
				m.cursorCellUp()
				So(*m.cursor, ShouldResemble, point{0, 0})
			})

			Convey("You can move right", func() {
				m.cursorCellRight()
				So(*m.cursor, ShouldResemble, point{1, 0})
			})

			Convey("You can move left", func() {
				m.cursorCellRight()
				m.cursorCellLeft()
				So(*m.cursor, ShouldResemble, point{0, 0})
			})

			Convey("You can't move up beyond the top", func() {
				m.cursorCellUp()
				So(*m.cursor, ShouldResemble, point{0, 0})
			})

			Convey("You can't move left beyond the edge", func() {
				m.cursorCellLeft()
				So(*m.cursor, ShouldResemble, point{0, 0})
			})

			Convey("You can't move down below the bottom", func() {
				m.cursor = &point{15, 15}
				m.cursorCellDown()
				So(*m.cursor, ShouldResemble, point{15, 15})
			})

			Convey("You can't move right beyond the edge", func() {
				m.cursor = &point{15, 15}
				m.cursorCellRight()
				So(*m.cursor, ShouldResemble, point{15, 15})
			})
		})

		Convey("When moving section to section", func() {
			Convey("You can move down", func() {
				m.cursorSectionDown()
				So(*m.cursor, ShouldResemble, point{0, 4})
			})

			Convey("You can move up", func() {
				m.cursorSectionDown()
				m.cursorSectionUp()
				So(*m.cursor, ShouldResemble, point{0, 0})
			})

			Convey("You can move right", func() {
				m.cursorSectionRight()
				So(*m.cursor, ShouldResemble, point{4, 0})
			})

			Convey("You can move left", func() {
				m.cursorSectionRight()
				m.cursorSectionLeft()
				So(*m.cursor, ShouldResemble, point{0, 0})
			})

			Convey("You can't move up beyond the top", func() {
				m.cursorCellUp()
				So(*m.cursor, ShouldResemble, point{0, 0})
			})

			Convey("You can't move left beyond the edge", func() {
				m.cursorCellLeft()
				So(*m.cursor, ShouldResemble, point{0, 0})
			})

			Convey("You can't move down below the bottom", func() {
				m.cursor = &point{15, 15}
				m.cursorSectionDown()
				So(*m.cursor, ShouldResemble, point{15, 15})
			})

			Convey("You can't move right beyond the edge", func() {
				m.cursor = &point{15, 15}
				m.cursorSectionRight()
				So(*m.cursor, ShouldResemble, point{15, 15})
			})
		})
	})
}
