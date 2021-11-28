package year2020

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/phinze/adventofcode/aoc"
)

func TestDayFour(t *testing.T) {
	in, err := aoc.FetchInput(2020, 4)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := DayFour(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("out: %s", spew.Sdump(out))
}
