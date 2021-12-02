package year2021

import (
	"testing"

	"github.com/phinze/adventofcode/aoc"
)

func TestDayTwo(t *testing.T) {
	in, err := aoc.FetchInput(2021, 2)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := DayTwo(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("out: %#v", out)
}
