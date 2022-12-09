package year2022

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/phinze/adventofcode/aoc"
)

func TestDayThree(t *testing.T) {
	in, err := aoc.FetchInput(2022, 3)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := DayThree(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("out: %s", spew.Sdump(out))
}