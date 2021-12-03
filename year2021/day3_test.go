package year2021

import (
	"testing"

	"github.com/phinze/adventofcode/aoc"
)

func TestDayThree(t *testing.T) {
	in, err := aoc.FetchInput(2021, 3)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := DayThree(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("out: %#v", out)
}
