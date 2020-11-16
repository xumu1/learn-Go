package demo

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestFirstTry(t *testing.T) {
	a := 1
	t.Log(a&Readable == Readable)
	t.Log(a&Writable == Writable)
	t.Log(a&Executable == Executable)
}
