package state

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestField(t *testing.T) {
	Convey("Given a field of cells", t, func() {
		p := Point{0, 0}
		f := newField(2)
		So(f.size, ShouldEqual, 2)
		So(len(f.cells), ShouldEqual, 4)

		Convey("You can load from a bytearray", func() {
			f := fieldFromBytes([]byte("\x00\x00\x00\x00"))
			So(f.size, ShouldEqual, 2)
			So(len(f.cells), ShouldEqual, 4)
		})

		Convey("You can get an index from a point", func() {
			So(f.i(Point{1, 1}), ShouldEqual, 3)
		})

		Convey("You can get a string representation of the field", func() {
			So(f.String(), ShouldEqual, "\x00\x00\x00\x00")
		})

		Convey("You can manage the state of a cell in the field", func() {
			So(f.state(p), ShouldBeFalse)
			f.vivify(p)
			So(f.state(p), ShouldBeTrue)
			f.kill(p)
			So(f.state(p), ShouldBeFalse)
		})

		Convey("You can manage the flags of a cell", func() {
			f.setCorrect(p, true)
			So(f.correct(p), ShouldBeTrue)
			f.setComplete(p, true)
			So(f.complete(p), ShouldBeTrue)
			f.mark(p)
			So(f.marked(p), ShouldBeTrue)
			f.flag(p)
			So(f.flagged(p), ShouldBeTrue)
			f.clear(p)
			So(f.correct(p), ShouldBeFalse)
			So(f.complete(p), ShouldBeFalse)
			So(f.marked(p), ShouldBeFalse)
			So(f.flagged(p), ShouldBeFalse)
		})
	})
}
