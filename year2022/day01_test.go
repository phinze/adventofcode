package year2022

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/phinze/adventofcode/aoc"
)

func TestDayOne(t *testing.T) {
	in, err := aoc.FetchInput(2022, 1)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := DayOne(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("out: %s", spew.Sdump(out))
}
