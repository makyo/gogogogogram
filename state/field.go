package state

import "math"

// field represents a square field of cells.
type field struct {
	cells []cell
	size  int
}

// newField returns a field of cells, all unset.
func newField(size int) *field {
	return &field{
		cells: make([]cell, size*size),
		size:  size,
	}
}

// fieldFromBytes returns a field of cells  given a bytearray; it assumes that the bytearray is a square.
func fieldFromBytes(repr []byte) *field {
	f := &field{
		cells: make([]cell, len(repr)),
		size:  int(math.Sqrt(float64(len(repr)))),
	}
	for i := range repr {
		f.cells[i] = cell(repr[i])
	}
	return f
}

// i is a utility function for translating a point to an array index
func (f *field) i(p Point) int {
	return p.Y*f.size + p.X
}

// String returns a string representation of the field.
func (f *field) String() string {
	return string(f.cells)
}

// state returns whether or not the cell is alive.
func (f *field) state(p Point) bool {
	return f.cells[f.i(p)].state()
}

// correct returns whether or not the guess for the cell's state is correct, or whether or not the section is made up of all complete guesses.
func (f *field) correct(p Point) bool {
	return f.cells[f.i(p)].correct()
}

// complete returns whether or not the current section is complete (that is, all sections in its row/column are correct).
func (f *field) complete(p Point) bool {
	return f.cells[f.i(p)].complete()
}

// flagged returns whether or not the cell is suspected of being dead.
func (f *field) flagged(p Point) bool {
	return f.cells[f.i(p)].flagged()
}

// marked returns whether or not the cell is suspected of being alive.
func (f *field) marked(p Point) bool {
	return f.cells[f.i(p)].marked()
}

// vivify sets the cell to alive.
func (f *field) vivify(p Point) {
	f.cells[f.i(p)] = f.cells[f.i(p)].vivify()
}

// kill sets the cell to dead.
func (f *field) kill(p Point) {
	f.cells[f.i(p)] = f.cells[f.i(p)].kill()
}

// setCorrect sets the whether or not the guess at the cell's state is correct
func (f *field) setCorrect(p Point, to bool) {
	f.cells[f.i(p)] = f.cells[f.i(p)].setCorrect(to)
}

// setComplete sets whether or not the section is complete (that is, all of its rows/columns are correct).
func (f *field) setComplete(p Point, to bool) {
	f.cells[f.i(p)] = f.cells[f.i(p)].setComplete(to)
}

// mark marks the cell as suspected alive.
func (f *field) mark(p Point) {
	f.cells[f.i(p)] = f.cells[f.i(p)].mark()
}

// flag marks the cell as suspected dead.
func (f *field) flag(p Point) {
	f.cells[f.i(p)] = f.cells[f.i(p)].flag()
}

// clear clears all but the cell's state.
func (f *field) clear(p Point, deadCorrect bool) {
	f.cells[f.i(p)] = f.cells[f.i(p)].clear(deadCorrect)
}
