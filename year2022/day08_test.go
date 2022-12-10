package year2022

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/phinze/adventofcode/aoc"
	"github.com/stretchr/testify/require"
)

func TestDayEight_small(t *testing.T) {
	in := `30373
25512
65332
33549
35390
`
	out, err := DayEight(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	require.Equal(t, 21, out.VisibleTrees)
	require.Equal(t, 8, out.MaxScenicScore)
}

func TestDayEight(t *testing.T) {
	in, err := aoc.FetchInput(2022, 8)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	out, err := DayEight(in)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("out: %s", spew.Sdump(out))
}
