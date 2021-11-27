package year2020

import (
	"testing"

	"github.com/phinze/adventofcode/aoc"
)

func TestDayOne(t *testing.T) {
	in, err := aoc.FetchInput(2020, 1)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := DayOne(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("out: %s", out)
}
