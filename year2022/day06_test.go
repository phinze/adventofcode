package year2022

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/phinze/adventofcode/aoc"
)

func TestDaySix(t *testing.T) {
	in, err := aoc.FetchInput(2022, 6)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := DaySix(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("out: %s", spew.Sdump(out))
}
