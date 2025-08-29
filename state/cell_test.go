package state

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCell(t *testing.T) {
	Convey("Given a cell", t, func() {
		var c cell

		Convey("When managing state", func() {
			So(c.state(), ShouldBeFalse)

			Convey("You can vivify and kill the cell", func() {
				So(c.state(), ShouldBeFalse)
				c = c.vivify()
				So(c.state(), ShouldBeTrue)
				c = c.kill()
				So(c.state(), ShouldBeFalse)
			})
		})

		Convey("When managing flags", func() {

			Convey("You can mark a cell as assumed alive", func() {
				So(c.marked(), ShouldBeFalse)
				c = c.mark()
				So(c.marked(), ShouldBeTrue)
			})

			Convey("You can mark a cell as assumed dead", func() {
				So(c.flagged(), ShouldBeFalse)
				c = c.flag()
				So(c.flagged(), ShouldBeTrue)
			})

			Convey("You can mark/unmark a cell as correct", func() {
				So(c.correct(), ShouldBeFalse)
				c = c.setCorrect(true)
				So(c.correct(), ShouldBeTrue)
				c = c.setCorrect(false)
				So(c.correct(), ShouldBeFalse)
			})

			Convey("You can mark/unmark a cell as complete", func() {
				So(c.complete(), ShouldBeFalse)
				c = c.setComplete(true)
				So(c.complete(), ShouldBeTrue)
				c = c.setComplete(false)
				So(c.complete(), ShouldBeFalse)
			})

			Convey("Marking a cell unflags it and vice versa", func() {
				c = c.mark()
				So(c.marked(), ShouldBeTrue)
				So(c.flagged(), ShouldBeFalse)
				c = c.flag()
				So(c.marked(), ShouldBeFalse)
				So(c.flagged(), ShouldBeTrue)
				c = c.mark()
				So(c.marked(), ShouldBeTrue)
				So(c.flagged(), ShouldBeFalse)
			})

			Convey("Marking/flagging a cell manages its correctness", func() {
				c = c.vivify()
				So(c.correct(), ShouldBeFalse)
				c = c.mark()
				So(c.correct(), ShouldBeTrue)
				c = c.kill()
				So(c.correct(), ShouldBeFalse)

				c = c.vivify()
				c = c.flag()
				So(c.correct(), ShouldBeFalse)
				c = c.kill()
				So(c.correct(), ShouldBeTrue)
			})

			Convey("You can clear all flags (but leave the state)", func() {
				c = c.vivify()
				c = c.mark()
				c = c.setCorrect(true)
				c = c.setComplete(true)
				So(c.state(), ShouldBeTrue)
				So(c.marked(), ShouldBeTrue)
				So(c.flagged(), ShouldBeFalse)
				So(c.correct(), ShouldBeTrue)
				So(c.complete(), ShouldBeTrue)
				c = c.clear(false)
				So(c.state(), ShouldBeTrue)
				So(c.marked(), ShouldBeFalse)
				So(c.flagged(), ShouldBeFalse)
				So(c.correct(), ShouldBeFalse)
				So(c.complete(), ShouldBeFalse)
			})
		})
	})
}
