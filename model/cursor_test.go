package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCursor(t *testing.T) {
	Convey("Given a cursor", t, func() {

		Convey("When moving cell to cell", func() {
			m, err := New(4, 4)
			So(err, ShouldBeNil)
			So(m.cursor, ShouldEqual, 0)

			Convey("You can move down", func() {
				m = m.CursorCellDown()
				So(m.cursor, ShouldEqual, 16)
			})

			Convey("You can move up", func() {
				m = m.CursorCellDown()
				m = m.CursorCellUp()
				So(m.cursor, ShouldEqual, 0)
			})

			Convey("You can move right", func() {
				m = m.CursorCellRight()
				So(m.cursor, ShouldEqual, 1)
			})

			Convey("You can move left", func() {
				m = m.CursorCellRight()
				m = m.CursorCellLeft()
				So(m.cursor, ShouldEqual, 0)
			})

			Convey("You can't move up beyond the top", func() {
				m = m.CursorCellUp()
				So(m.cursor, ShouldEqual, 0)
			})

			Convey("You can't move left beyond the edge", func() {
				m = m.CursorCellLeft()
				So(m.cursor, ShouldEqual, 0)
			})

			bottomRight := (m.size * m.size) - 1
			m.cursor = bottomRight

			Convey("You can't move down below the bottom", func() {
				m.cursor = bottomRight
				m = m.CursorCellDown()
				So(m.cursor, ShouldEqual, bottomRight)
			})

			Convey("You can't move right beyond the edge", func() {
				m.cursor = bottomRight
				m = m.CursorCellRight()
				So(m.cursor, ShouldEqual, bottomRight)
			})
		})

		Convey("When moving section to section", func() {
			m, err := New(4, 4)
			So(err, ShouldBeNil)
			Convey("You can move down", func() {
				m = m.CursorSectionDown()
				So(m.cursor, ShouldEqual, 64)
			})

			Convey("You can move up", func() {
				m = m.CursorSectionDown()
				m = m.CursorSectionUp()
				So(m.cursor, ShouldEqual, 0)
			})

			Convey("You can move right", func() {
				m = m.CursorSectionRight()
				So(m.cursor, ShouldEqual, 4)
			})

			Convey("You can move left", func() {
				m = m.CursorSectionRight()
				m = m.CursorSectionLeft()
				So(m.cursor, ShouldEqual, 0)
			})

			Convey("You can't move up beyond the top", func() {
				m = m.CursorCellUp()
				So(m.cursor, ShouldEqual, 0)
			})

			Convey("You can't move left beyond the edge", func() {
				m = m.CursorCellLeft()
				So(m.cursor, ShouldEqual, 0)
			})

			bottomRight := (m.size * m.size) - 1

			Convey("You can't move down below the bottom", func() {
				m.cursor = bottomRight
				m = m.CursorSectionDown()
				So(m.cursor, ShouldEqual, bottomRight)
			})

			Convey("You can't move right beyond the edge", func() {
				m.cursor = bottomRight
				m = m.CursorSectionRight()
				So(m.cursor, ShouldEqual, bottomRight)
			})
		})
	})
}
